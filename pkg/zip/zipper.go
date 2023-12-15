package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipDirectory zips the contents of the specified directory.
func ZipDirectory(srcDir, destZip string) error {
	zipFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

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

		file, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipFileWriter, file)
		return err
	})

	return err
}
