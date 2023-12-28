package plugin

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/hashicorp/go-plugin"
	"github.com/zcubbs/blueprint"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// DiscoverAndLoadBlueprintPlugins discovers, loads, and registers blueprint plugins
func DiscoverAndLoadBlueprintPlugins(pluginDir string) ([]blueprint.Generator, error) {
	var plugins []blueprint.Generator

	log.Debug("Discovering plugins in", "path", pluginDir)

	err := filepath.Walk(pluginDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isPluginFile(path) {
			generator, err := loadAndRegisterPlugin(path)
			if err != nil {
				return err
			}
			plugins = append(plugins, generator)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return plugins, nil
}

// loadAndRegisterPlugin handles the loading and registration of a single plugin
func loadAndRegisterPlugin(path string) (blueprint.Generator, error) {
	client, err := NewBlueprintPluginRPCClient(path)
	if err != nil {
		return nil, err
	}
	//defer client.Cleanup()

	generator, err := client.Dispense()
	if err != nil {
		return nil, fmt.Errorf("failed to dispense plugin %s: %v", path, err)
	}

	spec, err := generator.LoadSpec()
	if err != nil {
		return nil, fmt.Errorf("failed to load spec from plugin %s: %v", path, err)
	}

	err = blueprint.Register(blueprint.Blueprint{
		Spec:      spec,
		Generator: generator,
		Type:      blueprint.TypePlugin, // Or other appropriate type
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register blueprint from plugin %s: %v", path, err)
	}

	return generator, nil
}

// isPluginFile checks if the file at the given path is a plugin executable
func isPluginFile(path string) bool {
	var pluginExtension string

	switch runtime.GOOS {
	case "windows":
		pluginExtension = ".exe" // Windows expects .exe files for executables
	default:
		pluginExtension = "" // Linux and macOS do not use file extensions for executables
	}

	if pluginExtension == "" {
		return !strings.Contains(filepath.Base(path), ".") // No dot in the filename typically means an executable in Unix-like systems
	}
	return strings.HasSuffix(path, pluginExtension)
}

// RPCPluginClient handles client-side communication for blueprint plugins
type RPCPluginClient struct {
	client *plugin.Client
}

// NewBlueprintPluginRPCClient creates a new RPC client for a blueprint plugin
func NewBlueprintPluginRPCClient(pluginPath string) (*RPCPluginClient, error) {
	// Ensure the plugin path is compatible with the OS
	pluginPath = strings.ReplaceAll(pluginPath, "\\", "/")

	// Create the command to execute the plugin
	cmd := exec.Command(pluginPath)

	// Initialize the plugin client with the appropriate configuration
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: blueprint.HandshakeConfig,
		Plugins: map[string]plugin.Plugin{
			"blueprint": &blueprint.GeneratorPlugin{},
		},
		Cmd:              cmd,
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC},
	})

	return &RPCPluginClient{client: client}, nil
}

// Dispense returns the actual blueprint plugin implementation
func (b *RPCPluginClient) Dispense() (blueprint.Generator, error) {
	rpcClient, err := b.client.Client()
	if err != nil {
		return nil, err
	}

	raw, err := rpcClient.Dispense("blueprint")
	if err != nil {
		return nil, err
	}

	return raw.(blueprint.Generator), nil
}

// Cleanup should be called to clean up the client resources
func (b *RPCPluginClient) Cleanup() {
	b.client.Kill()
}
