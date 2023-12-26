package buildins_helloworld

import (
	_ "embed"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/blueprint"
	"os"
	"path/filepath"
	"text/template"
	// Additional imports as necessary
)

//go:embed spec.yaml
var specYaml []byte

// Generator conforms to the ComponentGenerator interface from the blueprint package
type Generator struct{}

// Generate implements the ComponentGenerator interface, generating a Go API server
func (g *Generator) Generate(spec blueprint.Spec, values map[string]string, workdir string) error {
	// Step 1: Parse ComponentSpec to get required configurations
	config, err := parseConfig(spec, values)
	if err != nil {
		return fmt.Errorf("error parsing component spec: %v", err)
	}

	// Step 2: Create project structure and files based on the parsed config
	projectPath, err := createProjectStructure(workdir, config)
	if err != nil {
		return fmt.Errorf("error creating project structure: %v", err)
	}

	// Step 3: Generate project files
	err = generateProjectFiles(projectPath, config)
	if err != nil {
		return fmt.Errorf("error generating project files: %v", err)
	}

	return nil
}

func (g *Generator) LoadSpec() (blueprint.Spec, error) {
	return blueprint.LoadBlueprintSpecFromBytes(specYaml)
}

// parseConfig extracts configuration options from ComponentSpec
func parseConfig(spec blueprint.Spec, values map[string]string) (Config, error) {
	var config Config

	// iterate over options and set config values
	// if no value is provided, use the default value
	for _, option := range spec.Options {
		switch option.ID {
		case "option1":
			if value, ok := values[option.ID]; ok {
				config.Option1 = value
			} else {
				config.Option1 = option.Default
			}
		}
	}

	return config, nil
}

// createProjectStructure sets up the project directory and base files
func createProjectStructure(outputPath string, _ Config) (string, error) {
	projectPath := filepath.Join(outputPath, "hello-world")
	if err := os.MkdirAll(projectPath, 0755); err != nil {
		return "", err
	}

	return projectPath, nil
}

// generateProjectFiles generates necessary Go files for the API server
func generateProjectFiles(projectPath string, config Config) error {
	// Define the file paths and corresponding templates
	files := map[string]string{
		"hello.txt": helloTxtTemplate,
	}

	// Process each template and create files
	for filePath, tmpl := range files {
		fullPath := filepath.Join(projectPath, filePath)
		if err := processTemplate(fullPath, tmpl, config); err != nil {
			return err
		}
	}

	return nil
}

func processTemplate(filePath, tmpl string, config Config) error {
	t, err := template.New("template").Parse(tmpl)
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error("Failed to close file", "package", "apiserver", "function", "processTemplate", "error", err)
		}
	}(file)

	return t.Execute(file, config)
}

// Config represents the configuration options for the API server
type Config struct {
	Option1 string
}
