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
	toPrint := make([]interface{}, 0, len(v)+1)
	toPrint = append(toPrint, "[DEBUG] ")
	toPrint = append(toPrint, v...)
	t.logger.Println(toPrint...)
}

func (t *TermLogger) Info(v ...interface{}) {
	toPrint := make([]interface{}, 0, len(v)+1)
	toPrint = append(toPrint, "[INFO] ")
	toPrint = append(toPrint, v...)
	t.logger.Println(toPrint...)
}

func (t *TermLogger) Warn(v ...interface{}) {
	toPrint := make([]interface{}, 0, len(v)+1)
	toPrint = append(toPrint, "[WARN] ")
	toPrint = append(toPrint, v...)
	t.logger.Println(toPrint...)
}

func (t *TermLogger) Error(v ...interface{}) {
	toPrint := make([]interface{}, 0, len(v)+1)
	toPrint = append(toPrint, "[ERROR] ")
	toPrint = append(toPrint, v...)
	t.logger.Println(toPrint...)
}

func (t *TermLogger) Fatal(v ...interface{}) {
	toPrint := make([]interface{}, 0, len(v)+1)
	toPrint = append(toPrint, "[FATAL] ")
	toPrint = append(toPrint, v...)
	t.logger.Fatalln(toPrint...)
}
