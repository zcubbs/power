package designer

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/pkg/blueprint"
	"github.com/zcubbs/power/pkg/blueprint/go/apiserver"
	"github.com/zcubbs/power/pkg/zip"
	"os"
	"path"
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
	// todo: implement this

	return nil
}

// validateOptionType checks if the provided option type aligns with the blueprint spec.
func validateOptionType(option blueprint.Option) error {
	// todo: implement this

	return nil
}

// validateOptionChoices checks if the provided option choices aligns with the blueprint spec.
func validateOptionChoices(option blueprint.Option) error {
	// todo: implement this

	return nil
}

// EnableBuiltinGenerators Register Built-in Generators
func EnableBuiltinGenerators() error {
	if err := apiserver.Register(); err != nil {
		return err
	}

	return nil
}
