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
