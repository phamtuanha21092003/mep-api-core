package logger

import (
	"os"

	"github.com/phamtuanha21092003/mep-api-core/pkg/config"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var logger = &Logger{}

// SetUpLogger settings
func SetUpLogger() {
	logger = &Logger{logrus.New()}
	logger.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		}}
	logger.SetOutput(os.Stdout)
	if config.AppCfg().Debug {
		logger.SetLevel(logrus.DebugLevel)
	}
}

// GetLogger return default logger
func GetLogger() *Logger {
	return logger
}
