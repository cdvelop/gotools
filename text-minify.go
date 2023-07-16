package gotools

import "strings"

func TextMinify(text *string) {

	if text != nil {

		// Eliminar espacios en blanco al inicio y al final
		*text = strings.TrimSpace(*text)

		// Eliminar espacios en blanco entre caracteres y saltos de línea
		*text = strings.Join(strings.Fields(*text), "")

		// Convertir a minúsculas
		*text = strings.ToLower(*text)

	}

}
