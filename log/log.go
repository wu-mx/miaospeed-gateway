package log

import (
	"gitlab.com/CoiaPrant/clog"
)

var Logger clog.Logger = clog.Child("", clog.LevelInfo)

func Print(a ...interface{}) {
	Logger.Print(a...)
}

func Info(a ...interface{}) {
	Logger.Info(a...)
}

func Success(a ...interface{}) {
	Logger.Success(a...)
}

func Warn(a ...interface{}) {
	Logger.Warn(a...)
}

func Error(a ...interface{}) {
	Logger.Error(a...)
}

func Fatal(a ...interface{}) {
	Logger.Fatal(a...)
}

func Debug(a ...interface{}) {
	Logger.Debug(a...)
}

func Printf(format string, a ...interface{}) {
	Logger.Printf(format, a...)
}

func Infof(format string, a ...interface{}) {
	Logger.Infof(format, a...)
}

func Successf(format string, a ...interface{}) {
	Logger.Successf(format, a...)
}

func Warnf(format string, a ...interface{}) {
	Logger.Warnf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	Logger.Errorf(format, a...)
}

func Fatalf(format string, a ...interface{}) {
	Logger.Fatalf(format, a...)
}

func Debugf(format string, a ...interface{}) {
	Logger.Debugf(format, a...)
}

func IsDebug() bool {
	return Logger.IsDebug()
}

func SetDebug(enabled bool) {
	Logger.SetDebug(enabled)
}

func DebugFlag() *bool {
	return Logger.DebugFlag()
}
