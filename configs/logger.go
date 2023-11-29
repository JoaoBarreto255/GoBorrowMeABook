package configs


import (
	"log"
	"io"
	"os"
	"strings"
)

type Logger struct {
	debug *log.Logger
	info *log.Logger
	warning *log.Logger
	err *log.Logger
	writer io.Writer
}

func NewLogger() *Logger {
	writer := io.Writer(os.Stdout)
	flags := log.Ldate|log.Ltime
	
	return &Logger{
		debug: log.New(writer, "DEBUG: ", flags),
		info: log.New(writer, "INFO: ", flags),
		warning: log.New(writer, "WARNING: ", flags),
		err: log.New(writer, "ERROR: ", flags),
		writer: writer
	}
}

func (l *Logger) Log(level string, v ... interface{}) {
	level = strings.ToLower(level)
	switch level {
	case "info":
		l.info.Println(v ...)
	case "warning":
		l.warning.Println(v ...)
	case "warn"
		l.warning.Println(v ...)
	case "error":
		l.err.Println(v ...)
	default:
		l.debug.Println(v ...)
	}
}

func (l *Logger) Debug(v ... interface{}) {
	l.debug.Println(v ...)
}

func (l *Logger) Info(v ... interface{}) {
	l.info.Println(v ...)
}

func (l *Logger) Warn(v ... interface{}) {
	l.warning.Println(v ...)
}

func (l *Logger) Error(v ... interface{}) {
	l.err.Println(v ...)
}