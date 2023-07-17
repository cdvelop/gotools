package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func TestAddStringContendFromDirAndExtension(t *testing.T) {

	var css string
	gotools.AddStringContentFromFile("test/style.css", &css)

	var file_txt string
	gotools.AddStringContentFromFile("test/file.txt", &file_txt)
	gotools.AddStringContentFromFile("test/inputs_1.txt", &file_txt)

	var testData = map[string]struct {
		dir    string
		ext    string
		out    string
		expect string
	}{
		"leyendo fichero único css":           {"test", ".css", "", css},
		"leyendo 2 ficheros txt":              {"test", ".txt", "", file_txt},
		"extension no existe no retorna nada": {"test", ".nil", "", ""},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			err := gotools.AddStringContendFromDirAndExtension(data.dir, data.ext, &data.out)

			if data.out != data.expect {
				log.Fatalf("Para la entrada '%s', '%s' se esperaba '%v' pero se obtuvo '%v'\nerr: %v", data.dir, data.ext, data.expect, data.out, err)
			}
		})
	}
}
