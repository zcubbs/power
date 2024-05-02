package pluginloader

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"github.com/zcubbs/blueprint"
	"github.com/zcubbs/power/pkg/plugin"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// LoadPluginsFromArchive loads plugins from a given archive (.tar.gz or .zip) and then uses
// DiscoverAndLoadBlueprintPlugins to load and register them.
func LoadPluginsFromArchive(archivePath string) ([]blueprint.Generator, error) {
	// Determine the file type (tar.gz or zip) based on the file extension
	if strings.HasSuffix(archivePath, ".tar.gz") {
		return loadFromTarGz(archivePath)
	} else if strings.HasSuffix(archivePath, ".zip") {
		return loadFromZip(archivePath)
	}

	return nil, fmt.Errorf("unsupported archive format: %s", archivePath)
}

// loadFromTarGz extracts tar.gz files to a temporary directory and returns the path.
func loadFromTarGz(archivePath string) ([]blueprint.Generator, error) {
	archivePath = filepath.Clean(archivePath)
	file, err := os.Open(archivePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}
	defer gzr.Close()

	tarReader := tar.NewReader(gzr)

	tempDir, err := os.MkdirTemp("", "plugins")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir) // Clean up

	if err := extractTarFiles(tarReader, tempDir); err != nil {
		return nil, err
	}

	return plugin.DiscoverAndLoadBlueprintPlugins(tempDir)
}

// extractTarFiles extracts files from a tar reader to a target directory.
func extractTarFiles(tarReader *tar.Reader, targetDir string) error {
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			return nil // End of archive
		}
		if err != nil {
			return err
		}

		cleanPath, err := SanitizeArchivePath(targetDir, header.Name)
		if err != nil {
			return err
		}

		if err := createTarEntry(header, tarReader, cleanPath); err != nil {
			return err
		}
	}
}

// createTarEntry creates a file or directory from a tar header.
func createTarEntry(header *tar.Header, tarReader *tar.Reader, targetPath string) error {
	switch header.Typeflag {
	case tar.TypeDir:
		return os.MkdirAll(targetPath, 0750)
	case tar.TypeReg:
		return writeFileFromTar(tarReader, targetPath)
	default:
		return fmt.Errorf("unsupported file type in tar: %v", header.Typeflag)
	}
}

// writeFileFromTar writes a file from tar data to the given path.
func writeFileFromTar(tarReader *tar.Reader, fileP string) error {
	fileP = filepath.Clean(fileP)
	outFile, err := os.Create(fileP)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, tarReader)
	return err
}

// loadFromZip extracts zip files to a temporary directory and returns the path.
func loadFromZip(archivePath string) ([]blueprint.Generator, error) {
	// Open the zip file
	zipFile, err := zip.OpenReader(archivePath)
	if err != nil {
		return nil, err
	}
	defer zipFile.Close()

	// Create a temporary directory to extract files
	tempDir, err := os.MkdirTemp("", "plugins")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir) // Clean up

	// Extract files to the temporary directory
	for _, file := range zipFile.File {
		// Path for the extracted file or directory
		fPath, err := SanitizeArchivePath(tempDir, file.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to sanitize archive path: %w", err)
		}

		// Check for ZipSlip (Directory traversal vulnerability)
		if !strings.HasPrefix(fPath, filepath.Clean(tempDir)+string(os.PathSeparator)) {
			return nil, fmt.Errorf("illegal file path: %s", fPath)
		}

		if file.FileInfo().IsDir() {
			// Create directory
			if err := os.MkdirAll(fPath, os.ModePerm); err != nil {
				return nil, err
			}
		} else {
			// Create file
			if err := extractFile(file, fPath); err != nil {
				return nil, err
			}
		}
	}

	// Load and register plugins from the temporary directory
	return plugin.DiscoverAndLoadBlueprintPlugins(tempDir)
}

// extractFile extracts a single file from a zip archive
func extractFile(file *zip.File, fPath string) error {
	rc, err := file.Open()
	if err != nil {
		return err
	}
	defer rc.Close()

	// clean the file path
	fPath = filepath.Clean(fPath)
	outFile, err := os.OpenFile(fPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}

	_, err = io.CopyN(outFile, rc, int64(file.UncompressedSize64))
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	err = outFile.Close()
	if err != nil {
		return fmt.Errorf("failed to close file: %w", err)
	} // Close the file immediately on error or after copy
	return err
}

// SanitizeArchivePath Sanitize archive file pathing from "G305: Zip Slip vulnerability"
func SanitizeArchivePath(d, t string) (v string, err error) {
	v = filepath.Join(d, t)
	if strings.HasPrefix(v, filepath.Clean(d)) {
		return v, nil
	}

	return "", fmt.Errorf("%s: %s", "content filepath is tainted", t)
}
