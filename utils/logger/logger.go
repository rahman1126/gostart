package logger

import log "github.com/sirupsen/logrus"

func SetJSONLogger() {
	formatter := new(log.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(formatter)
}

func Trace(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Trace(args)
}

func Info(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Info(args)
}

func Debug(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Debug(args)
}

func Warn(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Warn(args)
}

func Error(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Error(args)
}

func Panic(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Panic(args)
}

func Fatal(fields map[string]interface{}, args ...interface{}) {
	log.WithFields(fields).Fatal(args)
}