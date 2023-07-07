package gotools_test

import (
	"testing"

	"github.com/cdvelop/gotools"
)

func TestFindFilesWithNonZeroSize(t *testing.T) {
	dir := "."                                   // Directorio actual
	filenames := []string{"README.md", "go.mod"} // Archivos que estamos buscando

	err := gotools.FindFilesWithNonZeroSize(dir, filenames)
	if err != nil {
		t.Errorf("Se produjo un error inesperado: %v", err)
	}
}
