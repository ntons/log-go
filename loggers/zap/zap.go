package zap

import (
	"go.uber.org/zap"

	"github.com/ntons/log-go"
)

type zapLogger struct {
	sugar *zap.SugaredLogger
}

func New(l *zap.Logger, skip int) log.Logger {
	return &zapLogger{
		l.WithOptions(zap.AddCallerSkip(skip)).Sugar(),
	}
}

func (r zapLogger) Debug(args ...interface{}) {
	r.sugar.Debug(args...)
}
func (r zapLogger) Info(args ...interface{}) {
	r.sugar.Info(args...)
}
func (r zapLogger) Warn(args ...interface{}) {
	r.sugar.Warn(args...)
}
func (r zapLogger) Error(args ...interface{}) {
	r.sugar.Error(args...)
}
func (r zapLogger) Panic(args ...interface{}) {
	r.sugar.Panic(args...)
}
func (r zapLogger) Fatal(args ...interface{}) {
	r.sugar.Fatal(args...)
}

func (r zapLogger) Debugf(format string, args ...interface{}) {
	r.sugar.Debugf(format, args...)
}
func (r zapLogger) Infof(format string, args ...interface{}) {
	r.sugar.Infof(format, args...)
}
func (r zapLogger) Warnf(format string, args ...interface{}) {
	r.sugar.Warnf(format, args...)
}
func (r zapLogger) Errorf(format string, args ...interface{}) {
	r.sugar.Errorf(format, args...)
}
func (r zapLogger) Panicf(format string, args ...interface{}) {
	r.sugar.Panicf(format, args...)
}
func (r zapLogger) Fatalf(format string, args ...interface{}) {
	r.sugar.Fatalf(format, args...)
}

func (r zapLogger) Debugw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Debugw(msg, keyValuePairs...)
}
func (r zapLogger) Infow(msg string, keyValuePairs ...interface{}) {
	r.sugar.Infow(msg, keyValuePairs...)
}
func (r zapLogger) Warnw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Warnw(msg, keyValuePairs...)
}
func (r zapLogger) Errorw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Errorw(msg, keyValuePairs...)
}
func (r zapLogger) Panicw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Panicw(msg, keyValuePairs...)
}
func (r zapLogger) Fatalw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Fatalw(msg, keyValuePairs...)
}

func (r zapLogger) With(keyValuePairs ...interface{}) log.Logger {
	return zapLogger{sugar: r.sugar.With(keyValuePairs...)}
}
