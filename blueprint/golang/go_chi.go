package goblueprint

import "github.com/zcubbs/power/blueprint"

type GoChiGenerator struct{}

func (g *GoChiGenerator) Generate(spec blueprint.ComponentSpec, outputPath string) error {
	// Implementation for generating a project using go-chi
	// This can include setting up the main.go, configuring routes, etc.
	return nil
}

func init() {
	blueprint.RegisterGenerator("go-chi", &GoChiGenerator{})
}
