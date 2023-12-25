package designer

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/pkg/blueprint"
	"github.com/zcubbs/power/pkg/blueprint/go/apiserver"
	"github.com/zcubbs/power/pkg/zip"
	"os"
	"path"
	"strconv"
	"time"
)

// PostGenHook is a function that is called after a project is generated.
type PostGenHook func(archivePath string) error

// Generate iterates over components and generates each part.
func Generate(blueprintId string, values map[string]string, postGenHook PostGenHook) error {
	// get blueprint spec
	spec, err := blueprint.GetBlueprintSpec(blueprintId)
	if err != nil {
		return fmt.Errorf("error getting blueprint spec: %v", err)
	}

	timestamp := time.Now().Format("20060102150405")

	// directory & archive paths
	tmpDir := path.Join(os.TempDir(), "power", blueprintId, timestamp)
	workdir := path.Join(tmpDir, "workdir")

	// create tmp dir
	if err := os.MkdirAll(workdir, 0750); err != nil {
		return fmt.Errorf("error creating tmp dir: %v", err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			log.Error("Failed to remove temp dir",
				"package", "designer",
				"function", "Generate",
				"path", path,
				"error", err)
		}
	}(tmpDir)

	// load generator
	generator, specExists := blueprint.GetGenerator(blueprintId)
	if !specExists {
		return fmt.Errorf("no generator found for type: %s", spec.ID)
	}

	// validate spec config & values
	if err := validateSpecConfig(*spec, values); err != nil {
		return fmt.Errorf("validation error for '%s': %v", spec.ID, err)
	}

	// generate dir & file structure for blueprint & values
	if err := generator.Generate(*spec, values, workdir); err != nil {
		return fmt.Errorf("error generating component '%s': %v", spec.ID, err)
	}

	// create archive
	archivePath, err := createArchive(fmt.Sprintf("%s-%s", spec.ID, timestamp), "zip", tmpDir, workdir)
	if err != nil {
		return fmt.Errorf("error creating archive: %v", err)
	}

	// call post gen hook
	if postGenHook != nil {
		if err := postGenHook(archivePath); err != nil {
			return fmt.Errorf("error executing post gen hook: %v", err)
		}
	}

	return nil
}

// createArchive creates an archive of the generated project.
func createArchive(archiveName, format, tmpDir, workdir string) (string, error) {
	archivePath := path.Join(tmpDir, fmt.Sprintf("%s.%s", archiveName, format))
	if err := zip.Directory(workdir, archivePath); err != nil {
		return "", fmt.Errorf("error creating archive: %v", err)
	}
	return archivePath, nil
}

// validateComponentConfig checks if the provided component configuration
// aligns with the blueprint spec.
func validateSpecConfig(spec blueprint.Spec, values map[string]string) error {
	// validate options
	for _, option := range spec.Options {
		if err := validateOption(option); err != nil {
			return err
		}

		if err := validateOptionValue(option, values); err != nil {
			return err
		}
	}

	return nil
}

// validateOption checks if the provided option aligns with the blueprint spec.
func validateOption(option blueprint.Option) error {
	if err := validateOptionType(option); err != nil {
		return err
	}

	if err := validateOptionChoices(option); err != nil {
		return err
	}

	return nil
}

// validateOptionValue checks if the provided option value aligns with the blueprint spec.
func validateOptionValue(option blueprint.Option, values map[string]string) error {
	value, ok := values[option.ID]
	if !ok {
		return fmt.Errorf("value for option '%s' is missing", option.Name)
	}

	switch option.Type {
	case "text":
		return validateTextOption(option, value)
	case "number":
		return validateNumberOption(option, value)
	case "boolean":
		return validateBooleanOption(option, value)
	case "select":
		return validateSelectOption(option, value)
	}

	return nil
}

func validateTextOption(option blueprint.Option, value string) error {
	if value == "" {
		return fmt.Errorf("empty value for text option '%s'", option.Name)
	}
	return nil
}

func validateNumberOption(option blueprint.Option, value string) error {
	if _, err := strconv.ParseFloat(value, 64); err != nil {
		return fmt.Errorf("invalid number '%s' for option '%s'", value, option.Name)
	}
	return nil
}

func validateBooleanOption(option blueprint.Option, value string) error {
	if value != "true" && value != "false" {
		return fmt.Errorf("invalid boolean '%s' for option '%s'", value, option.Name)
	}
	return nil
}

func validateSelectOption(option blueprint.Option, value string) error {
	for _, choice := range option.Choices {
		if value == choice {
			return nil
		}
	}
	return fmt.Errorf("invalid choice '%s' for select option '%s'", value, option.Name)
}

// validateOptionType checks if the provided option type aligns with the blueprint spec.
func validateOptionType(option blueprint.Option) error {
	allowedTypes := map[string]bool{
		"boolean": true,
		"text":    true,
		"select":  true,
		"number":  true,
	}

	if _, ok := allowedTypes[option.Type]; !ok {
		return fmt.Errorf("invalid type '%s' for option '%s'", option.Type, option.Name)
	}

	return nil
}

// validateOptionChoices checks if the provided option choices aligns with the blueprint spec.
func validateOptionChoices(option blueprint.Option) error {
	// This validation is relevant only for options of type 'select'
	if option.Type != "select" {
		return nil
	}

	// Check if there are choices provided for the select type
	if len(option.Choices) == 0 {
		return fmt.Errorf("no choices provided for select option '%s'", option.Name)
	}

	// check for duplicate choices or any other specific rules
	choiceSet := make(map[string]bool)
	for _, choice := range option.Choices {
		if _, exists := choiceSet[choice]; exists {
			return fmt.Errorf("duplicate choice '%s' found in option '%s'", choice, option.Name)
		}
		choiceSet[choice] = true
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
