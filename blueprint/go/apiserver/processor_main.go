package apiserver

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/blueprint"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateMainGo(outputPath string, spec blueprint.ComponentSpec) error {
	// load template from bytes
	tpl, err := template.New("main.go").Parse(MainFileTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	// sanitize paths
	mainGoPath := filepath.Join(outputPath, "main.go")
	mainGoPath = filepath.Join(filepath.Clean(mainGoPath))
	file, err := os.Create(mainGoPath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Failed to close file",
				"package", "apiserver",
				"function", "GenerateMainGo",
				"error", err,
				"path", mainGoPath,
			)
		}
	}(file)

	return tpl.Execute(file, spec)
}
