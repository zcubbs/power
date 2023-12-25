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
		p, err := plugin.Open(file)
		if err != nil {
			return fmt.Errorf("failed to open plugin: %v", err)
		}

		symbol, err := p.Lookup("Plugin")
		if err != nil {
			return fmt.Errorf("failed to lookup symbol Plugin: %v", err)
		}

		bp, ok := symbol.(Generator)
		if !ok {
			return fmt.Errorf("invalid plugin type. expected Generator, got %T", symbol)
		}

		specFilePath := filepath.Join(pluginDir, "spec.yaml")
		spec, err := LoadBlueprintSpec(specFilePath)
		if err != nil {
			return fmt.Errorf("failed to load spec for plugin. make sure spec.yaml exists in the plugin directory: %v", err)
		}

		err = Register(Blueprint{
			Spec:      spec,
			Generator: bp,
		})
		if err != nil {
			return fmt.Errorf("failed to register plugin: %v", err)
		}
	}

	return nil
}
