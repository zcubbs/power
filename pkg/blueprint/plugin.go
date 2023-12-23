package blueprint

import (
	"log"
	"path/filepath"
	"plugin"
)

func LoadPlugins(pluginDir string) {
	files, err := filepath.Glob(filepath.Join(pluginDir, "*.so"))
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		p, err := plugin.Open(file)
		if err != nil {
			log.Println("Failed to load plugin:", err)
			continue
		}

		symbol, err := p.Lookup("Plugin")
		if err != nil {
			log.Println("Failed to find 'Plugin' symbol:", err)
			continue
		}

		bp, ok := symbol.(ComponentGenerator)
		if !ok {
			log.Println("Invalid plugin type:", file)
			continue
		}

		blueprintName := extractBlueprintName(file)
		err = RegisterGenerator(blueprintName, bp)
		if err != nil {
			log.Println("Failed to register plugin:", err)
			continue
		}

		specFilePath := filepath.Join(pluginDir, blueprintName+"-spec.yaml")
		spec, err := LoadBlueprintSpec(specFilePath)
		if err != nil {
			log.Println("Failed to load spec for plugin:", err)
			continue
		}
		RegisterBlueprintSpec(blueprintName, spec)
	}
}

func extractBlueprintName(pluginFilePath string) string {
	// Implement logic to extract the blueprint name
	return filepath.Base(pluginFilePath)
}
