package zip

import (
	"archive/zip"
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
				"error", err,
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
				"error", err,
				"path", destZip,
			)
		}
	}(zipWriter)

	err = filepath.Walk(srcDir, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // Skip directories
		}

		relPath, err := filepath.Rel(srcDir, filePath)
		if err != nil {
			return err
		}

		zipFileWriter, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		// sanitize paths
		filePath = filepath.Join(filepath.Clean(filePath))
		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			log.Error("Failed to close file",
				"package", "zip",
				"function", "Directory",
				"error", file.Close(),
				"path", filePath,
			)
		}(file)

		_, err = io.Copy(zipFileWriter, file)
		return err
	})

	return err
}
