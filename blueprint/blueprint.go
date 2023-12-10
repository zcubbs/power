package blueprint

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
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

// BlueprintSpec represents the structure of a YAML blueprint spec.
type BlueprintSpec struct {
	Name        string            `yaml:"name"`
	Description string            `yaml:"description"`
	Options     []BlueprintOption `yaml:"options"`
}

type BlueprintOption struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Choices     []string `yaml:"choices,omitempty"`
}

// Generators is a map that holds all registered blueprint generators.
var Generators sync.Map

// blueprintSpecs holds the specs for each registered blueprint.
var blueprintSpecs sync.Map

// RegisterGenerator adds a new generator to the system.
func RegisterGenerator(name string, generator ComponentGenerator) error {
	if _, loaded := Generators.LoadOrStore(name, generator); loaded {
		return fmt.Errorf("generator already registered: %s", name)
	}
	return nil
}

// RegisterBlueprintSpec registers the spec for a given blueprint type.
func RegisterBlueprintSpec(blueprintType string, spec *BlueprintSpec) {
	blueprintSpecs.Store(blueprintType, spec)
}

// GetGenerator retrieves a generator from the registry.
func GetGenerator(name string) (ComponentGenerator, bool) {
	generator, ok := Generators.Load(name)
	if !ok {
		return nil, false
	}
	return generator.(ComponentGenerator), true
}

// GetBlueprintSpec retrieves the spec for a given blueprint type.
func GetBlueprintSpec(blueprintType string) (*BlueprintSpec, error) {
	spec, ok := blueprintSpecs.Load(blueprintType)
	if !ok {
		return nil, fmt.Errorf("no spec found for blueprint type: %s", blueprintType)
	}

	blueprintSpec, ok := spec.(*BlueprintSpec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type for blueprint type: %s", blueprintType)
	}

	return blueprintSpec, nil
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

// LoadBlueprintSpec reads and parses the YAML spec file for a blueprint.
func LoadBlueprintSpec(filePath string) (*BlueprintSpec, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var spec BlueprintSpec
	err = yaml.Unmarshal(data, &spec)
	return &spec, err
}
