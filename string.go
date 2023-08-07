package gotools

import (
	"fmt"
	"strings"
)

func ExtractTwoPointArgument(option string, field *string) error {
	parts := strings.Split(option, ":")
	if len(parts) >= 2 {
		*field = strings.Join(parts[1:], ":")
	} else {
		return fmt.Errorf("delimitador ':' no encontrado en la cadena %v", option)
	}
	return nil
}

// remover \cmd ej: "mi_directorio\cmd"  "\\cmd"
func RemoveSuffixIfPresent(path *string, suffix string) {
	if path != nil {
		if strings.Contains(*path, suffix) {
			// Si el path tiene la extensión que queremos eliminar, la quitamos
			*path = (*path)[:len(*path)-len(suffix)]
		}
	}
}
