package telego

import (
	"log"
	"os"
)

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
}

type logger struct {
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (logger *logger) Debug(v ...interface{}) {
	logger.debugLogger.Println(v...)
}
func (logger *logger) Info(v ...interface{}) {
	logger.infoLogger.Println(v...)
}
func (logger *logger) Warn(v ...interface{}) {
	logger.warnLogger.Println(v...)
}
func (logger *logger) Error(v ...interface{}) {
	logger.errorLogger.Println(v...)
}

func getInitalLogger() Logger {
	flags := log.LstdFlags
	debugLogger := log.New(os.Stdout, "Telego Debug: ", flags)
	infoLogger := log.New(os.Stdout, "Telego Info: ", flags)
	warnLogger := log.New(os.Stdout, "Telego WARN: ", flags)
	errorLogger := log.New(os.Stdout, "Telego ERROR: ", flags)
	return &logger{
		debugLogger: debugLogger,
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
}
