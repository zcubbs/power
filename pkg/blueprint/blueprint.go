package blueprint

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"sync"
)

// Generator defines the interface for a blueprint.
type Generator interface {
	Generate(spec Spec, values map[string]string, workdir string) error
}

// Spec BlueprintSpec represents the structure of a YAML blueprint spec.
type Spec struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Options     []Option `yaml:"options"`
}

func (b *Spec) String() string {
	return fmt.Sprintf("Name: %s, Description: %s", b.Name, b.Description)
}

type Option struct {
	ID          string   `yaml:"id"`
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
var blueprints sync.Map

type Blueprint struct {
	Spec      *Spec
	Generator Generator
}

// Register registers the spec for a given blueprint type.
func Register(blueprint Blueprint) error {
	if blueprint.Spec.ID == "" {
		return fmt.Errorf("failed to register blueprint spec: missing id")
	}
	blueprints.Store(blueprint.Spec.ID, blueprint.Spec)
	Generators.Store(blueprint.Spec.ID, blueprint.Generator)

	return nil
}

// GetGenerator retrieves a generator from the registry.
func GetGenerator(id string) (Generator, bool) {
	generator, ok := Generators.Load(id)
	if !ok {
		return nil, false
	}
	return generator.(Generator), true
}

// GetBlueprintSpec retrieves the spec for a given blueprint type.
func GetBlueprintSpec(id string) (*Spec, error) {
	spec, ok := blueprints.Load(id)
	if !ok {
		return nil, fmt.Errorf("no spec found for blueprint id: %s", id)
	}

	blueprintSpec, ok := spec.(*Spec)
	if !ok {
		return nil, fmt.Errorf("invalid spec type for blueprint id: %s", id)
	}

	return blueprintSpec, nil
}

// LoadBlueprintSpec reads and parses the YAML spec file for a blueprint.
func LoadBlueprintSpec(filePath string) (*Spec, error) {
	// sanitize paths
	filePath = filepath.Join(filepath.Clean(filePath))
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
	blueprints.Range(func(key, value interface{}) bool {
		specs[key.(string)] = value.(*Spec)
		return true
	})
	return specs
}
