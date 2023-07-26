package gotools

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteIfFolderSizeExceeds(folder_path string, maxSizeMB int64) error {
	var size int64

	err := filepath.Walk(folder_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	if err != nil {
		return err
	}

	sizeMB := size / 1024 / 1024
	if sizeMB > maxSizeMB {
		err := os.RemoveAll(folder_path)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateFolderIfNotExist(folder_path string) error {
	// Verificar si el directorio ya existe
	_, err := os.Stat(folder_path)
	if err == nil {
		// El directorio ya existe, no es necesario crearlo
		return nil
	}

	// Intentar crear el directorio
	err = os.MkdirAll(folder_path, os.ModePerm)
	if err != nil {
		// Ocurrió un error al crear el directorio
		return err
	}

	// Directorio creado exitosamente
	return nil
}

const DeleteEmptyFolderResponse = "directorio con archivos. no fue eliminado"

func DeleteEmptyFolder(folder_path string) error {
	// fmt.Println("Verificar si el directorio existe")
	fileInfo, err := os.Stat(folder_path)
	if err != nil {
		// log.Println(err)
		// fmt.Println("directorio no existe nada que eliminar")
		return nil
	}

	// fmt.Println("Verificar si es un directorio")
	if !fileInfo.IsDir() {
		return os.ErrInvalid
	}

	// fmt.Println("Obtener la lista de archivos y directorios en el directorio")
	fileList, err := os.ReadDir(folder_path)
	if err != nil {
		return err
	}

	// fmt.Println("Verificar si el directorio está vacío")
	if len(fileList) == 0 {
		// fmt.Println("Eliminar el directorio vacío")
		err := os.Remove(folder_path)
		if err != nil {
			return err
		}
	} else {
		return fmt.Errorf(DeleteEmptyFolderResponse)
	}

	return nil
}
