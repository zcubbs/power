package zip

import (
	"archive/zip"
	"fmt"
	"github.com/charmbracelet/log"
	"io"
	"os"
	"path/filepath"
)

// Directory zips the contents of the specified directory.
func Directory(srcDir, destZip string) error {
	// sanitize paths
	destZip = filepath.Join(filepath.Clean(destZip))
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			log.Error("Failed to close zip file",
				"package", "zip",
				"function", "Directory",
				"error", err.Error(),
				"path", destZip,
			)
		}
	}(zipFile)

	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			log.Error("Failed to close zip writer",
				"package", "zip",
				"function", "Directory",
				"error", err.Error(),
				"path", destZip,
			)
		}
	}(zipWriter)

	err = filepath.Walk(srcDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(filePath) == ".zip" {
			return nil // Skip directories and zip files
		}

		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %w", err)
		}

		zipFileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("failed to create zip file writer: %w", err)
		}

		// sanitize paths
		filePath = filepath.Join(filepath.Clean(filePath))
		file, err := os.Open(filePath)
		if err != nil {
			return fmt.Errorf("failed to open file: %w", err)
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Error("Failed to close file",
					"package", "zip",
					"function", "Directory",
					"error", err,
					"path", filePath,
				)
			}
		}(file)

		_, err = io.Copy(zipFileWriter, file)
		return err
	})

	return err
}
