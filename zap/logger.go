package zap

import (
	"github.com/ntons/log-go"
	"go.uber.org/zap"
)

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

func (r Recorder) Debugw(msg string, keyValuePairs ...interface{}) {
	r.s.Debugw(msg, keyValuePairs...)
}
func (r Recorder) Infow(msg string, keyValuePairs ...interface{}) {
	r.s.Infow(msg, keyValuePairs...)
}
func (r Recorder) Warnw(msg string, keyValuePairs ...interface{}) {
	r.s.Warnw(msg, keyValuePairs...)
}
func (r Recorder) Errorw(msg string, keyValuePairs ...interface{}) {
	r.s.Errorw(msg, keyValuePairs...)
}
func (r Recorder) Panicw(msg string, keyValuePairs ...interface{}) {
	r.s.Panicw(msg, keyValuePairs...)
}
func (r Recorder) Fatalw(msg string, keyValuePairs ...interface{}) {
	r.s.Fatalw(msg, keyValuePairs...)
}

func (r Recorder) With(fields log.Fields) log.Recorder {
	zapfields := make([]interface{}, 0, len(fields)*2)
	for key, val := range fields {
		zapfields = append(zapfields, key, val)
	}
	return Recorder{s: r.s.With(zapfields...)}
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
