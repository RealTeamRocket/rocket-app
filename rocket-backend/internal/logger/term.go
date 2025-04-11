package logger

import (
	"log"
	"os"
)

type TermLogger struct {
	logger *log.Logger
}

func NewTermLogger() *TermLogger {
	return &TermLogger{
		logger: log.New(os.Stdout, "", log.LstdFlags),
	}
}

func (t *TermLogger) Debug(v ...interface{}) {
	t.logger.SetPrefix("DEBUG: ")
	t.logger.Println(v...)
}

func (t *TermLogger) Info(v ...interface{}) {
	t.logger.SetPrefix("INFO: ")
	t.logger.Println(v...)
}

func (t *TermLogger) Warn(v ...interface{}) {
	t.logger.SetPrefix("WARN: ")
	t.logger.Println(v...)
}

func (t *TermLogger) Error(v ...interface{}) {
	t.logger.SetPrefix("ERROR: ")
	t.logger.Println(v...)
}

func (t *TermLogger) Fatal(v ...interface{}) {
	t.logger.SetPrefix("FATAL: ")
	t.logger.Fatalln(v...)
}
