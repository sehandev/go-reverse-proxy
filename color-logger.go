package main

import (
	"fmt"
	"log"
)

// colorLogger comment
type colorLogger struct {
	debugStart   string
	infoStart    string
	warningStart string
	errorStart   string
	end          string
	checker      accessChecker
}

// newColorLogger construct
func newColorLogger() *colorLogger {
	logger := new(colorLogger)
	logger.debugStart = "\033[94m"
	logger.infoStart = "\033[92m"
	logger.warningStart = "\033[93m"
	logger.errorStart = "\033[91m"
	logger.end = "\033[0m"
	return logger
}

func (logger *colorLogger) debug(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s", logger.debugStart, message, logger.end)
	logger.checker.newLog(message)
}

func (logger *colorLogger) info(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s", logger.infoStart, message, logger.end)
	logger.checker.newLog(message)
}

func (logger *colorLogger) warning(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s", logger.warningStart, message, logger.end)
	logger.checker.newLog(message)
}

func (logger *colorLogger) error(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	log.Printf("%s%s%s", logger.errorStart, message, logger.end)
	logger.checker.newLog(message)
}

func (logger *colorLogger) test(format string, a ...interface{}) {
	logger.debug(format, a...)
	logger.info(format, a...)
	logger.warning(format, a...)
	logger.error(format, a...)
}
