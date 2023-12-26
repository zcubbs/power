package blueprint

import (
	"fmt"
	"path/filepath"
	"plugin"
)

func LoadPlugins(pluginDir string) error {
	files, err := filepath.Glob(filepath.Join(pluginDir, "*.so"))
	if err != nil {
		return fmt.Errorf("failed to list plugins: %v", err)
	}

	for _, file := range files {
		err := LoadPlugin(file)
		if err != nil {
			return fmt.Errorf("failed to load plugin %s: %v", file, err)
		}
	}

	return nil
}

func LoadPlugin(path string) error {
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

	// Assert the type to blueprint.Generator (or the correct interface)
	gen, ok := sym.(Generator)
	if !ok {
		return fmt.Errorf("unexpected type from module symbol")
	}

	// Load the spec
	spec, err := LoadBlueprintSpec(filepath.Join(path, "spec.yaml"))
	if err != nil {
		return fmt.Errorf("error loading spec: %v", err)
	}

	// Use gen as needed
	// For example, if you have a registration function in your application:
	err = Register(Blueprint{
		Spec:      spec,
		Generator: gen,
	})
	if err != nil {
		return fmt.Errorf("error registering blueprint generator: %v", err)
	}

	return nil
}
