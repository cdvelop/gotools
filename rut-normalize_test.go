package gotools_test

import (
	"log"
	"testing"

	"github.com/cdvelop/gotools"
)

func Test_RutNormalize(t *testing.T) {

	var (
		dataRuNormalize = map[string]struct {
			inputData string
			expected  string
		}{
			"cero agregado al inicio?":         {"07863697-1", "7863697-1"},
			"1 dígito verificador incorrecto?": {"7863697-0", "7863697-1"},
			"1 cero agregado al inicio?":       {"05335341-K", "5335341-k"},

			"dígito verificador incorrecto?": {"20373221-0", "20373221-k"},

			"sin espacio en blanco?":     {" ", " "},
			"nada de espacio en blanco?": {"    ", ""},
		}
	)

	for prueba, data := range dataRuNormalize {
		t.Run((prueba), func(t *testing.T) {
			err := gotools.RutNormalize(&data.inputData)
			var resp string
			if err != "" {
				resp = err
			}

			if data.inputData != data.expected {
				log.Println(prueba)
				log.Fatalf("resultado: [%v] expectativa: [%v]\nerror:[%v]\n%v", data.inputData, data.expected, resp, data.inputData)
			}
			// t.Logf("data entrada: [%v] resultado: [%v]\n", data.inputData, data.expected)
		})
	}
}
