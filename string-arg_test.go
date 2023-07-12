package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func TestGetPackage(t *testing.T) {
	var testData = map[string]struct {
		line   string
		expect string
		out    string
	}{
		"se espera x":                     {"valor:x", "x", ""},
		"se espera juan":                  {"user:juan", "juan", ""},
		"se espera error sin delimitador": {"userjuan", "", ""},
		"se espera error sin data":        {"", "", ""},
		"doble : se espera midato:es":     {"dato:midato:es", "midato:es", ""},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {
			err := gotools.ExtractTwoPointArgument(data.line, &data.out)

			if data.expect != data.out {
				t.Fatalf("Para la entrada '%s', se esperaba '%s' pero se obtuvo '%s' error: %v", data.line, data.expect, data.out, err)
				log.Fatal()
			}
		})
	}
}
