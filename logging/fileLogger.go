package logging

import (
	"fmt"
	"os"
)

type FileLogger struct {
	BaseLogger
	fileName string
	file     *os.File
}

func NewFileLogger(fileName string, utc bool) (*FileLogger, error) {
	if fileName == "" {
		return nil, fmt.Errorf("filename cannot be empty string")
	}

	fl := &FileLogger{
		fileName:   fileName,
		BaseLogger: BaseLogger{Utc: utc},
	}
	err := fl.open()
	if err != nil {
		return nil, err
	}

	return fl, nil
}

func (fl *FileLogger) open() error {
	file, err := os.OpenFile(fl.fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	fl.file = file
	return nil
}

func (fl *FileLogger) Close() {
	if fl.file != nil {
		fl.file.Close()
	}
}

func (fl *FileLogger) Info(msg string) {

	logMessage := fmt.Sprintf("%s : INFO : %s\n", fl.timestamp(), msg)
	_, e := fl.file.WriteString(logMessage)
	if e != nil {
		fmt.Printf("Error writing to log file: %s\n", e)
	}
}

func (fl *FileLogger) Error(msg string, err error) {

	logMessage := ""
	if err != nil {
		logMessage = fmt.Sprintf("%s : ERROR : %s; error - %s\n", fl.timestamp(), msg, err)
	} else {
		logMessage = fmt.Sprintf("%s : ERROR : %s\n", fl.timestamp(), msg)
	}
	_, e := fl.file.WriteString(logMessage)
	if e != nil {
		fmt.Printf("Error writing to log file: %s\n", e)
	}
}
