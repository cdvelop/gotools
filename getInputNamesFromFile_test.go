package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
	"github.com/cdvelop/model"
)

func TestGetInputNamesFromGoFile(t *testing.T) {
	var testData = map[string]struct {
		file_path string
		pattern   string
		expected  []string
	}{
		"búsqueda de inputs_1 se esperan 3": {"test/inputs_1.txt", model.INPUT_PATTERN, []string{"Phone", "RadioGenero", "TextArea"}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := gotools.GetNamesFromFileAndPattern(data.file_path, data.pattern)

			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v\nerr: %v", data.expected[i], valor, err)
				}
			}

		})
	}
}

func TestGetNamesFromDirectoryExtensionAndPattern(t *testing.T) {
	var testData = map[string]struct {
		file_path string
		extension string
		pattern   string
		expected  []string
	}{
		"búsqueda en 2 archivos goo, Phone repetido": {"test", ".goo", model.INPUT_PATTERN, []string{"Phone", "RadioGenero", "TextArea"}},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp, err := gotools.GetNamesFromDirectoryExtensionAndPattern(data.file_path, data.extension, data.pattern)

			// Comparar elemento por elemento
			for i, valor := range resp {
				if valor != data.expected[i] {
					log.Fatalf("se esperaba %v, pero se obtuvo %v\nresp:%v\nerr: %v", data.expected[i], valor, resp, err)
				}
			}

		})
	}
}
