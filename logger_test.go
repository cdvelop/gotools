package gotools_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	. "github.com/cdvelop/gotools"
)

func TestLogger(t *testing.T) {
	file_name := "log_test"

	log_text := "escribiendo Log %v"
	total := 10

	os.Remove(file_name + ".log")

	SetupLogsToFile(file_name)

	for i := 1; i < total; i++ {
		log.Printf(log_text, i)
	}

	expected := fmt.Sprintf(log_text, total-1)

	res := TextExists(file_name+".log", expected)
	if res != 1 {
		log.Fatalf("se esperaba como resultado 1 y se obtuvo: %v texto: %v", res, expected)
	}
}
