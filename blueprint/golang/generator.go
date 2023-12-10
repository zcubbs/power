package goblueprint

import (
	"fmt"
	"github.com/zcubbs/power/blueprint"
	"os"
	"path/filepath"
	"text/template"
)

type goGenerator struct{}

func (g *goGenerator) Generate(spec blueprint.ComponentSpec, outputPath string) error {
	routerOption, dbOption := spec.Config["router"].(string), spec.Config["database"].(string)

	if err := g.generateMainFile(routerOption, dbOption, outputPath); err != nil {
		return err
	}

	// Additional file generation as required

	return nil
}

func (g *goGenerator) generateMainFile(router, db, outputPath string) error {
	mainFilePath := filepath.Join(outputPath, "main.go")
	file, err := os.Create(mainFilePath)
	if err != nil {
		return fmt.Errorf("failed to create main file: %w", err)
	}
	defer file.Close()

	data := struct {
		RouterImport string
	}{
		RouterImport: g.routerImportPath(router),
	}

	tmpl, err := template.New("main").Parse(MainFileTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	if err := tmpl.Execute(file, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	return nil
}

func (g *goGenerator) routerImportPath(router string) string {
	switch router {
	case "go-chi":
		return "github.com/go-chi/chi"
	case "standard-lib":
		return "net/http"
	case "go-fiber":
		return "github.com/gofiber/fiber/v2"
	case "gorilla-mux":
		return "github.com/gorilla/mux"
	default:
		return "net/http" // default to standard library
	}
}

func init() {
	err := blueprint.RegisterGenerator("go", &goGenerator{})
	if err != nil {
		panic(err)
	}
}
