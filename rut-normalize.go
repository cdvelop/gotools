package gotools

import (
	"fmt"
	"strings"

	"github.com/cdvelop/input"
)

func RutNormalize(rutIn *string) (err string) {
	if rutIn != nil && len(*rutIn) > 3 {

		r := input.Rut()

		err = r.ValidateField(*rutIn, false)
		if err != "" {

			//limpiamos elementos en blanco
			NormalizeTextData(rutIn)
			neWRutIn := strings.ToLower(*rutIn)

			if len(neWRutIn) < 2 {
				return "run menor a 2 dígitos"
			}
			neWRutIn = neWRutIn[1:] //primera prueba quitando el primer carácter
			err = r.ValidateField(neWRutIn, false)
			if err == "" {
				*rutIn = neWRutIn
			} else { //cambiemos el dígito verificador y probemos
				_, onlyRun, _ := input.RunData(*rutIn)
				dv := input.DvRut(onlyRun)
				run2 := fmt.Sprintf("%v-%v", onlyRun, dv)
				// fmt.Printf("NUEVO RUN: %v\n", run2)
				err = r.ValidateField(run2, false) //ultima opción

				if err == "" {
					*rutIn = strings.ToLower(run2)
				} else {
					return "no se logro normalizar run " + *rutIn
				}
			}
		}
	} else {
		return "tamaño rut inválido"
	}
	return ""
}
