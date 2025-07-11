package logger

import (
	logger "github.com/sirupsen/logrus"
)

func LogInfo(logStr interface{}) {
	logger.Info(logStr)
}

func LogWarning(logStr interface{}) {
	logger.Warning(logStr)
}

func LogError(logStr interface{}) {
	logger.Error(logStr)
}

func LogDebug(logStr interface{}) {
	logger.Debug(logStr)
}

func LogTrace(logStr interface{}) {
	logger.Trace(logStr)
}
