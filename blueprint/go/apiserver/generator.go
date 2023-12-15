package apiserver

import (
	_ "embed"
	"github.com/zcubbs/power/blueprint"
	"github.com/zcubbs/power/pkg/zip"
	"os"
	"path/filepath"
)

type Generator struct{}

func (g *Generator) Generate(spec blueprint.ComponentSpec, outputPath string) error {
	// Create the output directory if it doesn't exist
	err := os.MkdirAll(outputPath, 0755)
	if err != nil {
		return err
	}

	// Generate main.go from the template
	if err := GenerateMainGo(outputPath, spec); err != nil {
		return err
	}

	// Define the output zip file path
	zipFilePath := filepath.Join(outputPath, "project.zip")

	// Zip the contents of the output directory
	return zip.ZipDirectory(outputPath, zipFilePath)
}

//go:embed spec.yaml
var specFS []byte

func Register() error {
	spec, err := blueprint.LoadBlueprintSpecFromBytes(specFS)
	if err != nil {
		return err
	}
	blueprint.RegisterBlueprintSpec("go-api-server", spec)
	return blueprint.RegisterGenerator("go-api-server", &Generator{})
}
