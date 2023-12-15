package apiserver

import (
	"github.com/zcubbs/power/blueprint"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateMainGo(outputPath string, spec blueprint.ComponentSpec) error {
	// load template from bytes
	tpl, err := template.New("main.go").Parse(MainFileTemplate)

	mainGoPath := filepath.Join(outputPath, "main.go")
	file, err := os.Create(mainGoPath)
	if err != nil {
		return err
	}
	defer file.Close()

	return tpl.Execute(file, spec)
}
