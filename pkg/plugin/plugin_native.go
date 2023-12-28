package plugin

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/blueprint"
	"path/filepath"
	"plugin"
	"reflect"
)

// LoadNativePlugins loads all plugins in the given directory
// Deprecated: Use DiscoverAndLoadBlueprintPlugins instead
func LoadNativePlugins(pluginDir string) error {
	files, err := filepath.Glob(filepath.Join(pluginDir, "*.so"))
	if err != nil {
		return fmt.Errorf("failed to list plugins: %v", err)
	}

	for _, file := range files {
		err := LoadNativePlugin(file)
		if err != nil {
			return fmt.Errorf("failed to load plugin %s: %v", file, err)
		}
	}

	return nil
}

func LoadNativePlugin(path string) error {
	// Open the plugin
	plug, err := plugin.Open(path)
	if err != nil {
		return fmt.Errorf("error opening plugin %s: %v", path, err)
	}

	// Look up the exported symbol
	sym, err := plug.Lookup("Plugin")
	if err != nil {
		return fmt.Errorf("error looking up symbol 'Plugin' in %s: %v", path, err)
	}

	log.Debug("Loaded plugin", "path", path, "sym", reflect.TypeOf(sym))

	// Assert the type to blueprint.Generator (or the correct interface)
	gen, ok := sym.(*blueprint.Generator)
	if !ok {
		return fmt.Errorf("unexpected type from module symbol")
	}

	// gen to *blueprint.Generator
	dgen := *gen

	// Load the spec
	spec, err := dgen.LoadSpec()
	if err != nil {
		return fmt.Errorf("error loading spec: %v", err)
	}

	// Use gen as needed
	// For example, if you have a registration function in your application:
	err = blueprint.Register(blueprint.Blueprint{
		Spec:      spec,
		Generator: dgen,
		Type:      blueprint.TypePlugin,
	})
	if err != nil {
		return fmt.Errorf("error registering blueprint generator: %v", err)
	}

	return nil
}
