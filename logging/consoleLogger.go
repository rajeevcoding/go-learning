package logging

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
}

func (logger *ConsoleLogger) Info(msg string) {
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	fmt.Printf("%s : INFO : %s\n", timestamp, msg)
}

func (logger *ConsoleLogger) Error(msg string, err error) {
	timestamp := time.Now().UTC().Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Printf("%s : ERROR : %s; error - %s\n", timestamp, msg, err)
	} else {
		fmt.Printf("%s : ERROR : %s\n", timestamp, msg)
	}
}
