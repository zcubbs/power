package blueprint

import (
	"log"
	"path/filepath"
	"plugin"
)

type Plugin interface {
	Generate(spec ComponentSpec, outputPath string) error
}

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

		bp, ok := symbol.(Plugin)
		if !ok {
			log.Println("Invalid plugin type:", file)
			continue
		}

		err = RegisterGenerator(file, bp)
		if err != nil {
			log.Println("Failed to register plugin:", err)
			continue
		}
	}
}
