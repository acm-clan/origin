package logger

import (
	"github.com/sirupsen/logrus"
)

// InitLogger init logger
func InitLogger(level string) {
	switch level {
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "error":
		logrus.SetLevel(logrus.ErrorLevel)
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	default:
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000000000",
	})
}

// Infof log info
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Info log info
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Errorf log error
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Error log error
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Debugf log debug
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Debug log debug
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Warnf log warn
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Warn log warn
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}
