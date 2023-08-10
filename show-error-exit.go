package gotools

import (
	"fmt"
	"os"
)

func ShowErrorAndExit(message interface{}) {
	PrintError(fmt.Sprintf("Error: %v\n", message))
	fmt.Println("")
	PrintInfo("Presiona enter para salir...\n")
	fmt.Scanln()
	os.Exit(1)
}
