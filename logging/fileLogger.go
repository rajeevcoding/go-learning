package logging

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	fileName string
	file     *os.File
}

func NewFileLogger(fileName string) (*FileLogger, error) {
	if fileName == "" {
		return nil, fmt.Errorf("filename cannot be empty string")
	}

	fl := &FileLogger{fileName: fileName}
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

	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("%s : INFO : %s\n", timestamp, msg)
	_, e := fl.file.WriteString(logMessage)
	if e != nil {
		fmt.Printf("Error writing to log file: %s\n", e)
	}
}

func (fl *FileLogger) Error(msg string, err error) {

	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	logMessage := ""
	if err != nil {
		logMessage = fmt.Sprintf("%s : ERROR : %s; error - %s\n", timestamp, msg, err)
	} else {
		logMessage = fmt.Sprintf("%s : ERROR : %s\n", timestamp, msg)
	}
	_, e := fl.file.WriteString(logMessage)
	if e != nil {
		fmt.Printf("Error writing to log file: %s\n", e)
	}
}
