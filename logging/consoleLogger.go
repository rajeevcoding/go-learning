package logging

import (
	"fmt"
)

type ConsoleLogger struct {
	BaseLogger
}

func (logger *ConsoleLogger) Info(msg string) {
	fmt.Printf("%s : INFO : %s\n", logger.timestamp(), msg)
}

func (logger *ConsoleLogger) Error(msg string, err error) {
	if err != nil {
		fmt.Printf("%s : ERROR : %s; error - %s\n", logger.timestamp(), msg, err)
	} else {
		fmt.Printf("%s : ERROR : %s\n", logger.timestamp(), msg)
	}
}
