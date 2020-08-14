package log

import (
	"errors"

	"go.uber.org/zap"
)

type zRecorder struct {
	sugar *zap.SugaredLogger
}

func (r zRecorder) Debug(args ...interface{}) {
	r.sugar.Debug(args...)
}
func (r zRecorder) Info(args ...interface{}) {
	r.sugar.Info(args...)
}
func (r zRecorder) Warn(args ...interface{}) {
	r.sugar.Warn(args...)
}
func (r zRecorder) Error(args ...interface{}) {
	r.sugar.Error(args...)
}
func (r zRecorder) Panic(args ...interface{}) {
	r.sugar.Panic(args...)
}
func (r zRecorder) Fatal(args ...interface{}) {
	r.sugar.Fatal(args...)
}

func (r zRecorder) Debugf(format string, args ...interface{}) {
	r.sugar.Debugf(format, args...)
}
func (r zRecorder) Infof(format string, args ...interface{}) {
	r.sugar.Infof(format, args...)
}
func (r zRecorder) Warnf(format string, args ...interface{}) {
	r.sugar.Warnf(format, args...)
}
func (r zRecorder) Errorf(format string, args ...interface{}) {
	r.sugar.Errorf(format, args...)
}
func (r zRecorder) Panicf(format string, args ...interface{}) {
	r.sugar.Panicf(format, args...)
}
func (r zRecorder) Fatalf(format string, args ...interface{}) {
	r.sugar.Fatalf(format, args...)
}

func (r zRecorder) Debugw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Debugw(msg, keyValuePairs...)
}
func (r zRecorder) Infow(msg string, keyValuePairs ...interface{}) {
	r.sugar.Infow(msg, keyValuePairs...)
}
func (r zRecorder) Warnw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Warnw(msg, keyValuePairs...)
}
func (r zRecorder) Errorw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Errorw(msg, keyValuePairs...)
}
func (r zRecorder) Panicw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Panicw(msg, keyValuePairs...)
}
func (r zRecorder) Fatalw(msg string, keyValuePairs ...interface{}) {
	r.sugar.Fatalw(msg, keyValuePairs...)
}

func (r zRecorder) With(fields Fields) Recorder {
	keyValuePairs := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyValuePairs = append(keyValuePairs, key, value)
	}
	return r.Withw(keyValuePairs...)
}
func (r zRecorder) Withw(keyValuePairs ...interface{}) Recorder {
	return zRecorder{sugar: r.sugar.With(keyValuePairs...)}
}

type zLogger struct{ *zap.SugaredLogger }

func newZapLogger(l *zap.Logger) Logger {
	if l == nil {
		panic(errors.New("zap logger is nil"))
	}
	return &zLogger{l.WithOptions(zap.AddCallerSkip(1)).Sugar()}
}

func SetZapLogger(l *zap.Logger) {
	var old Logger
	old, std = std, newZapLogger(l)
	old.Sync()
}

func (l zLogger) With(fields Fields) Recorder {
	keyValuePairs := make([]interface{}, 0, len(fields)*2)
	for key, value := range fields {
		keyValuePairs = append(keyValuePairs, key, value)
	}
	return l.Withw(keyValuePairs...)
}

func (l zLogger) Withw(keyValuePairs ...interface{}) Recorder {
	return &zRecorder{l.SugaredLogger.With(keyValuePairs...)}
}
