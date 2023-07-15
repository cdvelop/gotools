package gotools

import (
	"fmt"
	"os"
	"path/filepath"
)

// ej: gotools.DeleteFilesByExtension(main_folder\files, []string{".js", ".css", ".wasm"})
func DeleteFilesByExtension(dir string, exts []string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		for _, ext := range exts {
			if ext == filepath.Ext(file.Name()) {
				path := filepath.Join(dir, file.Name())
				err := os.Remove(path)
				if err != nil {
					fmt.Printf("Error deleting file: %s\n", err)
				}
				break
			}
		}
	}

	return nil
}
