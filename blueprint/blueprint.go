package blueprint

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
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

// Spec BlueprintSpec represents the structure of a YAML blueprint spec.
type Spec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Options     []Option `yaml:"options"`
}

func (b *Spec) String() string {
	return fmt.Sprintf("Name: %s, Description: %s", b.Name, b.Description)
}

type Option struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Type        string   `yaml:"type"`
	Default     string   `yaml:"default,omitempty"`
	Choices     []string `yaml:"choices,omitempty"`
}

func (b *Option) String() string {
	return fmt.Sprintf("Name: %s, Description: %s, Type: %s, Choices: %v", b.Name, b.Description, b.Type, b.Choices)
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
func RegisterBlueprintSpec(blueprintType string, spec *Spec) {
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
func GetBlueprintSpec(blueprintType string) (*Spec, error) {
	spec, ok := blueprintSpecs.Load(blueprintType)
	if !ok {
		return nil, fmt.Errorf("no spec found for blueprint type: %s", blueprintType)
	}

	blueprintSpec, ok := spec.(*Spec)
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
func LoadBlueprintSpec(filePath string) (*Spec, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var spec Spec
	err = yaml.Unmarshal(data, &spec)
	return &spec, err
}

// LoadBlueprintSpecFromBytes parses the YAML spec file for a blueprint from a byte array.
func LoadBlueprintSpecFromBytes(data []byte) (*Spec, error) {
	var spec Spec
	err := yaml.Unmarshal(data, &spec)
	return &spec, err
}

// GetAllBlueprintSpecs returns a map of all registered blueprint specs.
func GetAllBlueprintSpecs() map[string]*Spec {
	specs := make(map[string]*Spec)
	blueprintSpecs.Range(func(key, value interface{}) bool {
		specs[key.(string)] = value.(*Spec)
		return true
	})
	return specs
}
