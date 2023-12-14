package designer

import (
	"encoding/json"
	"fmt"
	"github.com/zcubbs/power/blueprint"
	"github.com/zcubbs/power/blueprint/go/apiserver"
)

// ProjectSpec represents the overall project specification.
type ProjectSpec struct {
	Components []blueprint.ComponentSpec `json:"components"`
}

// ParseSpec parses the JSON input into ProjectSpec struct.
func ParseSpec(jsonInput string) (*ProjectSpec, error) {
	var spec ProjectSpec
	err := json.Unmarshal([]byte(jsonInput), &spec)
	return &spec, err
}

// GenerateProject iterates over components and generates each part.
func GenerateProject(spec *ProjectSpec, outputPath string) error {
	for _, component := range spec.Components {
		generator, specExists := blueprint.GetGenerator(component.Type)
		if !specExists {
			return fmt.Errorf("no generator found for type: %s", component.Type)
		}

		// Retrieve the blueprint spec and validate the component config
		blueprintSpec, err := blueprint.GetBlueprintSpec(component.Type)
		if err != nil {
			return fmt.Errorf("error retrieving spec for '%s': %v", component.Type, err)
		}

		if err := validateComponentConfig(component, *blueprintSpec); err != nil {
			return fmt.Errorf("validation error for '%s': %v", component.Type, err)
		}

		if err := generator.Generate(component, outputPath); err != nil {
			return fmt.Errorf("error generating component '%s': %v", component.Type, err)
		}
	}

	return nil
}

// validateComponentConfig checks if the provided component configuration
// aligns with the blueprint spec.
func validateComponentConfig(component blueprint.ComponentSpec, spec blueprint.BlueprintSpec) error {
	for _, option := range spec.Options {
		value, ok := component.Config[option.Name]
		if !ok {
			return fmt.Errorf("required option '%s' not provided for '%s'", option.Name, component.Type)
		}

		if option.Type == "select" {
			validChoice := false
			for _, choice := range option.Choices {
				if value == choice {
					validChoice = true
					break
				}
			}
			if !validChoice {
				return fmt.Errorf("invalid choice for '%s': %v, must be one of %v", option.Name, value, option.Choices)
			}
		}
	}
	return nil
}

// EnableBuiltinGenerators Register Built-in Generators
func EnableBuiltinGenerators() error {
	if err := apiserver.Register(); err != nil {
		return err
	}

	return nil
}
