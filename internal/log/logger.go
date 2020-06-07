package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var logger *log.Entry

func Setup(version string) {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	logger = log.WithField("version", version)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
