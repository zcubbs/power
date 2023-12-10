package designer

import (
	"encoding/json"
	"github.com/zcubbs/power/blueprint"
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
		if err := blueprint.GenerateComponent(component, outputPath); err != nil {
			return err
		}
	}
	// Zip the project directory
	// ...
	return nil
}
