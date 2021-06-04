package zap

import (
	"go.uber.org/zap"

	log "github.com/ntons/log-go"
)

type zapLogger struct {
	sugar *zap.SugaredLogger
	level zap.AtomicLevel
}

func New(logger *zap.Logger, level zap.AtomicLevel) log.LoggerLevelSetter {
	return &zapLogger{
		sugar: logger.WithOptions(zap.AddCallerSkip(2)).Sugar(),
		level: level,
	}
}

func (z zapLogger) SetLevel(_level log.Level) {
	switch _level {
	case log.AllLevel:
		z.level.SetLevel(zap.DebugLevel)
	case log.DebugLevel:
		z.level.SetLevel(zap.DebugLevel)
	case log.InfoLevel:
		z.level.SetLevel(zap.InfoLevel)
	case log.WarnLevel:
		z.level.SetLevel(zap.WarnLevel)
	case log.ErrorLevel:
		z.level.SetLevel(zap.ErrorLevel)
	case log.PanicLevel:
		z.level.SetLevel(zap.PanicLevel)
	case log.FatalLevel:
		z.level.SetLevel(zap.FatalLevel)
	}
}

func (z zapLogger) Debug(args ...interface{}) {
	z.sugar.Debug(args...)
}
func (z zapLogger) Info(args ...interface{}) {
	z.sugar.Info(args...)
}
func (z zapLogger) Warn(args ...interface{}) {
	z.sugar.Warn(args...)
}
func (z zapLogger) Error(args ...interface{}) {
	z.sugar.Error(args...)
}
func (z zapLogger) Panic(args ...interface{}) {
	z.sugar.Panic(args...)
}
func (z zapLogger) Fatal(args ...interface{}) {
	z.sugar.Fatal(args...)
}

func (z zapLogger) Debugf(format string, args ...interface{}) {
	z.sugar.Debugf(format, args...)
}
func (z zapLogger) Infof(format string, args ...interface{}) {
	z.sugar.Infof(format, args...)
}
func (z zapLogger) Warnf(format string, args ...interface{}) {
	z.sugar.Warnf(format, args...)
}
func (z zapLogger) Errorf(format string, args ...interface{}) {
	z.sugar.Errorf(format, args...)
}
func (z zapLogger) Panicf(format string, args ...interface{}) {
	z.sugar.Panicf(format, args...)
}
func (z zapLogger) Fatalf(format string, args ...interface{}) {
	z.sugar.Fatalf(format, args...)
}

func (z zapLogger) Debugw(msg string, keyValuePairs ...interface{}) {
	z.sugar.Debugw(msg, keyValuePairs...)
}
func (z zapLogger) Infow(msg string, keyValuePairs ...interface{}) {
	z.sugar.Infow(msg, keyValuePairs...)
}
func (z zapLogger) Warnw(msg string, keyValuePairs ...interface{}) {
	z.sugar.Warnw(msg, keyValuePairs...)
}
func (z zapLogger) Errorw(msg string, keyValuePairs ...interface{}) {
	z.sugar.Errorw(msg, keyValuePairs...)
}
func (z zapLogger) Panicw(msg string, keyValuePairs ...interface{}) {
	z.sugar.Panicw(msg, keyValuePairs...)
}
func (z zapLogger) Fatalw(msg string, keyValuePairs ...interface{}) {
	z.sugar.Fatalw(msg, keyValuePairs...)
}

func (z zapLogger) With(keyValuePairs ...interface{}) log.Logger {
	return zapLogger{sugar: z.sugar.With(keyValuePairs...)}
}
