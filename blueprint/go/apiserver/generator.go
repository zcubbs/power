package apiserver

import (
	_ "embed"
	"github.com/zcubbs/power/blueprint"
)

type Generator struct{}

func (g *Generator) Generate(spec blueprint.ComponentSpec, outputPath string) error {
	// Implementation for generating a project
	return nil
}

//go:embed golang-spec.yaml
var specFS []byte

func Register() error {
	spec, err := blueprint.LoadBlueprintSpecFromBytes(specFS)
	if err != nil {
		return err
	}
	blueprint.RegisterBlueprintSpec("go-api-server", spec)
	return blueprint.RegisterGenerator("go-api-server", &Generator{})
}
