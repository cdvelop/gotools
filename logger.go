package gotools

import (
	"log"
	"os"
)

type logFile struct {
	file_name string
}

func SetupLogsToFile(log_file_name string) *logFile {
	nl := logFile{
		file_name: log_file_name,
	}
	log.SetOutput(&nl)
	return &nl
}

func (l logFile) Write(p []byte) (n int, err error) {

	f, err := os.OpenFile(l.file_name+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		PrintError("Archivo server.log no existe")
	}

	_, err = f.Write(p)
	if err != nil {
		PrintError("no se puede escribir en el archivo server.log")
	}

	return 0, nil
}
