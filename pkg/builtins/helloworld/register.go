package buildins_helloworld

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
		Type:      blueprint.TypeBuiltIn,
		Spec:      spec,
		Generator: &Generator{},
	})
}
