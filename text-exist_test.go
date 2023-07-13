package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func TestExist(t *testing.T) {
	var testData = map[string]struct {
		path   string
		search string
		expect int
	}{
		"archivo go.mod se espera 1":                    {"go.mod", "gotools", 1},
		"archivo go.mod contra contenido mismo archivo": {"go.mod", "go.mod", 1},
		"archivo file.txt en carpeta test":              {"test\\file.txt", "contenido", 1},
		"archivo file.txt contra archivo":               {"test\\file.txt", "test\\file.txt", 1},
		"sin data se espera 0":                          {"", "", 0},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			resp := gotools.TextExists(data.path, data.search)

			if resp != data.expect {
				t.Fatalf("Para la entrada '%s', '%s' se esperaba '%v' pero se obtuvo '%v'", data.path, data.search, data.expect, resp)
				log.Fatal()
			}
		})
	}
}
