package loggious

import (
	"os"

	"github.com/justjordant/cargo/internals/initUtils"

	"github.com/charmbracelet/log"
)

type Logger struct {
	*log.Logger
}

type LogType int

const (
	Info LogType = iota
	Fatal
	Error
	Warn
)

func New() *Logger {
	return &Logger{log.NewWithOptions(os.Stdout, log.Options{
		ReportCaller:    true,
		ReportTimestamp: true,
	})}
}

func (logger *Logger) logToFile(logType LogType, args ...interface{}) {
	file, err := os.OpenFile(initUtils.LogsPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.Fatal(err)
	}

	defer file.Close()

	logger.SetOutput(file)

	switch logType {
	case Info:
		logger.Info(args)
	case Fatal:
		logger.Fatal(args)
	case Error:
		logger.Error(args)
	case Warn:
		logger.Warn(args)
	}
}

func (logger *Logger) Log(logType LogType, args ...interface{}) {
	logger.logToFile(logType, args...)
}
