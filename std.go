package log

import (
	"go.uber.org/zap"
)

var std = func() Logger {
	logger, err := zap.NewDevelopment(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	return newZapLogger(logger)
}()

func Debug(args ...interface{}) {
	std.Debug(args...)
}
func Info(args ...interface{}) {
	std.Info(args...)
}
func Warn(args ...interface{}) {
	std.Warn(args...)
}
func Error(args ...interface{}) {
	std.Error(args...)
}
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Debugw(msg string, keyValuePairs ...interface{}) {
	std.Debugw(msg, keyValuePairs...)
}
func Infow(msg string, keyValuePairs ...interface{}) {
	std.Infow(msg, keyValuePairs...)
}
func Warnw(msg string, keyValuePairs ...interface{}) {
	std.Warnw(msg, keyValuePairs...)
}
func Errorw(msg string, keyValuePairs ...interface{}) {
	std.Errorw(msg, keyValuePairs...)
}
func Fatalw(msg string, keyValuePairs ...interface{}) {
	std.Fatalw(msg, keyValuePairs...)
}

func With(fields Fields) Recorder {
	return std.With(fields)
}
func Withw(keyValuePairs ...interface{}) Recorder {
	return std.Withw(keyValuePairs...)
}

func Sync() { std.Sync() }
