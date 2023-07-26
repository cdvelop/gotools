package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func TestDeleteEmptyFolder(t *testing.T) {
	var testData = map[string]struct {
		dir    string
		expect string
	}{
		"eliminación normal directorio sin data":                {"test/folder01", ""},
		"intento de eliminación de una carpeta con archivo":     {"test/folder", gotools.DeleteEmptyFolderResponse},
		"intento de eliminación de una carpeta con sub carpeta": {"test/folder0", gotools.DeleteEmptyFolderResponse},
	}

	for testName, data := range testData {
		t.Run(testName, func(t *testing.T) {

			err := gotools.CreateFolderIfNotExist(data.dir)
			if err != nil {
				log.Fatal(err)
			}

			err = gotools.DeleteEmptyFolder(data.dir)
			if err != nil {
				if data.expect != err.Error() {
					log.Fatalf("\nPara la entrada '%s'\n-se esperaba [%s]\n-se obtuvo [%s]", data.dir, data.expect, err.Error())
				}
			} else {
				if data.expect != "" {
					log.Fatalf("\nPara la entrada '%s'\n-se esperaba [%s]\n-no se obtuvo ningún error", data.dir, data.expect)
				}
			}

		})
	}
}
