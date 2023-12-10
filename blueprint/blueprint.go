package blueprint

import (
	"fmt"
	"sync"
)

// ComponentSpec is a generic specification for any project component.
type ComponentSpec struct {
	Type   string                 `json:"type"`
	Config map[string]interface{} `json:"config"`
}

// ComponentGenerator defines the interface for a blueprint.
type ComponentGenerator interface {
	Generate(spec ComponentSpec, outputPath string) error
}

// Generators is a map that holds all registered blueprint generators.
// Using a sync.Map for safe concurrent access.
var Generators sync.Map

// RegisterGenerator adds a new generator to the system.
func RegisterGenerator(name string, generator ComponentGenerator) error {
	if _, loaded := Generators.LoadOrStore(name, generator); loaded {
		return fmt.Errorf("generator already registered: %s", name)
	}
	return nil
}

// GetGenerator retrieves a generator from the registry.
func GetGenerator(name string) (ComponentGenerator, bool) {
	generator, ok := Generators.Load(name)
	if !ok {
		return nil, false
	}
	return generator.(ComponentGenerator), true
}

// GenerateComponent calls the appropriate generator based on the component type.
func GenerateComponent(spec ComponentSpec, outputPath string) error {
	value, ok := Generators.Load(spec.Type)
	if !ok {
		return fmt.Errorf("no generator found for type: %s", spec.Type)
	}

	generator, ok := value.(ComponentGenerator)
	if !ok {
		return fmt.Errorf("invalid generator for type: %s", spec.Type)
	}

	return generator.Generate(spec, outputPath)
}
