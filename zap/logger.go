package zap

import (
	"github.com/ntons/log-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func fitFields(fields log.Fields) []interface{} {
	arr := make([]interface{}, 0, len(fields)*2)
	for key, val := range fields {
		arr = append(arr, key, val)
	}
	return arr
}

func fitLevel(lev log.Level) zapcore.Level {
	switch lev {
	case log.DebugLevel:
		return zapcore.DebugLevel
	case log.InfoLevel:
		return zapcore.InfoLevel
	case log.WarnLevel:
		return zapcore.WarnLevel
	case log.ErrorLevel:
		return zapcore.ErrorLevel
	case log.PanicLevel:
		return zapcore.PanicLevel
	case log.FatalLevel:
		return zapcore.FatalLevel
	default:
		panic("invalid log level")
	}
}

type Recorder struct {
	s *zap.SugaredLogger
}

func (r Recorder) Debug(args ...interface{}) {
	r.s.Debug(args...)
}
func (r Recorder) Info(args ...interface{}) {
	r.s.Info(args...)
}
func (r Recorder) Warn(args ...interface{}) {
	r.s.Warn(args...)
}
func (r Recorder) Error(args ...interface{}) {
	r.s.Error(args...)
}
func (r Recorder) Panic(args ...interface{}) {
	r.s.Panic(args...)
}
func (r Recorder) Fatal(args ...interface{}) {
	r.s.Fatal(args...)
}

func (r Recorder) Debugf(format string, args ...interface{}) {
	r.s.Debugf(format, args...)
}
func (r Recorder) Infof(format string, args ...interface{}) {
	r.s.Infof(format, args...)
}
func (r Recorder) Warnf(format string, args ...interface{}) {
	r.s.Warnf(format, args...)
}
func (r Recorder) Errorf(format string, args ...interface{}) {
	r.s.Errorf(format, args...)
}
func (r Recorder) Panicf(format string, args ...interface{}) {
	r.s.Panicf(format, args...)
}
func (r Recorder) Fatalf(format string, args ...interface{}) {
	r.s.Fatalf(format, args...)
}

func (r Recorder) Debugw(msg string, fields log.Fields) {
	r.s.Debugw(msg, fitFields(fields)...)
}
func (r Recorder) Infow(msg string, fields log.Fields) {
	r.s.Infow(msg, fitFields(fields)...)
}
func (r Recorder) Warnw(msg string, fields log.Fields) {
	r.s.Warnw(msg, fitFields(fields)...)
}
func (r Recorder) Errorw(msg string, fields log.Fields) {
	r.s.Errorw(msg, fitFields(fields)...)
}
func (r Recorder) Panicw(msg string, fields log.Fields) {
	r.s.Panicw(msg, fitFields(fields)...)
}
func (r Recorder) Fatalw(msg string, fields log.Fields) {
	r.s.Fatalw(msg, fitFields(fields)...)
}

func (r Recorder) With(fields log.Fields) log.Recorder {
	return Recorder{s: r.s.With(fitFields(fields)...)}
}

type Logger struct {
	Recorder
}

func NewLogger(l *zap.Logger) Logger {
	// skip 2 more callstack, 1 for myself, 1 for log-go
	s := l.WithOptions(zap.AddCallerSkip(2)).Sugar()
	return Logger{Recorder: Recorder{s: s}}
}
func (l Logger) Close() error {
	return nil
}
func (l Logger) Sync() error {
	return l.s.Sync()
}
