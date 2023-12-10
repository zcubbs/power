package apiserver

import "github.com/zcubbs/power/blueprint"

type Generator struct{}

func (g *Generator) Generate(spec blueprint.ComponentSpec, outputPath string) error {
	// Implementation for generating a project
	return nil
}

func init() {
	err := blueprint.RegisterGenerator("go-api-server", &Generator{})
	if err != nil {
		panic(err)
	}
}
