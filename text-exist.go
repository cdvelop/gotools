package gotools

import (
	"os"
	"strings"
)

// ej: gotools.TextExists(index.html", "info_id")
// gotools.TextExists(static\main.js", "js\code.js")
func TextExists(file_path string, text_search_or_path_content string) int {
	file_content, err := os.ReadFile(file_path)
	if err != nil {
		return 0
	}
	string_content := string(file_content)
	TextMinify(&string_content)

	// Variable para almacenar el contenido de búsqueda
	var text_search string

	for _, value := range []string{".", "\\", "/"} {
		// Puede que sea una ruta
		if strings.Contains(text_search_or_path_content, value) {
			search_content, err := os.ReadFile(text_search_or_path_content)
			if err == nil {
				text_search = string(search_content)
			}
			break
		}
	}

	if text_search == "" {
		text_search = text_search_or_path_content
	}

	TextMinify(&text_search)

	// fmt.Printf("CONTENIDO: [%v]\n", string_content)
	// fmt.Printf("TEXTO A BUSCAR: [%v]\n", text_search)

	return strings.Count(string_content, text_search)
}
