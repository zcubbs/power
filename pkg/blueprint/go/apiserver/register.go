package apiserver

import (
	_ "embed"
	"github.com/zcubbs/power/pkg/blueprint"
)

//go:embed spec.yaml
var specFS []byte

func Register() error {
	spec, err := blueprint.LoadBlueprintSpecFromBytes(specFS)
	if err != nil {
		return err
	}
	return blueprint.Register(blueprint.Blueprint{
		Spec:      spec,
		Generator: &Generator{},
	})
}
