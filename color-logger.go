package main

import (
	"fmt"
	"log"
)

// ColorLogger comment
type ColorLogger struct {
	debugStart   string
	infoStart    string
	warningStart string
	errorStart   string
	end          string
}

// NewColorLogger construct
func NewColorLogger() *ColorLogger {
	logger := new(ColorLogger)
	logger.debugStart = "\033[94m"
	logger.infoStart = "\033[92m"
	logger.warningStart = "\033[93m"
	logger.errorStart = "\033[91m"
	logger.end = "\033[0m"
	return logger
}

func (logger *ColorLogger) debug(format string, a ...interface{}) {
	log.Printf("%s%s%s", logger.debugStart, fmt.Sprintf(format, a...), logger.end)
}

func (logger *ColorLogger) info(format string, a ...interface{}) {
	log.Printf("%s%s%s", logger.infoStart, fmt.Sprintf(format, a...), logger.end)
}

func (logger *ColorLogger) warning(format string, a ...interface{}) {
	log.Printf("%s%s%s", logger.warningStart, fmt.Sprintf(format, a...), logger.end)
}

func (logger *ColorLogger) error(format string, a ...interface{}) {
	log.Printf("%s%s%s", logger.errorStart, fmt.Sprintf(format, a...), logger.end)
}

func (logger *ColorLogger) test(format string, a ...interface{}) {
	logger.debug(format, a...)
	logger.info(format, a...)
	logger.warning(format, a...)
	logger.error(format, a...)
}
