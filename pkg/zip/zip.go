package zip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipFiles compresses one or more files into a single zip archive file.
// srcDir is the directory of the files to be zipped.
// destZip is the destination zip file path.
func ZipFiles(srcDir string, destZip string) error {
	outFile, err := os.Create(destZip)
	if err != nil {
		return err
	}
	defer outFile.Close()

	zipWriter := zip.NewWriter(outFile)
	defer zipWriter.Close()

	// Walk the directory and add files to the zip.
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

		zipFile, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fsFile.Close()

		_, err = io.Copy(zipFile, fsFile)
		return err
	})

	return err
}
