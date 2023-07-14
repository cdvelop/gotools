package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func TestGetInputNamesFromGoFile(t *testing.T) {
	var testData = map[string]struct {
		file_path string
		pattern   string
		expected  []string
	}{
		"búsqueda de inputs_1 se esperan 3": {"test/inputs_1.txt", `Input: input\.(\w+)\(\)`, []string{"Phone", "RadioGenero", "TextArea"}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp := gotools.GetNamesFromFileAndPattern(data.file_path, data.pattern)

			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v", data.expected[i], valor)
				}
			}

		})
	}
}
