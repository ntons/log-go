package logrus

import (
	"fmt"

	"github.com/ntons/log-go"
	"github.com/sirupsen/logrus"
)

func toFields(keyValuePairs []interface{}) logrus.Fields {
	fields := logrus.Fields{}
	for i := 1; i < len(keyValuePairs); i += 2 {
		key, ok := keyValuePairs[i-1].(string)
		if !ok {
			key = fmt.Sprintf("%v", keyValuePairs[i-1])
		}
		fields[key] = keyValuePairs[i]
	}
	return fields
}

type Recorder struct {
	e *logrus.Entry
}

func (r Recorder) Debug(args ...interface{}) {
	r.e.Debug(args...)
}
func (r Recorder) Info(args ...interface{}) {
	r.e.Info(args...)
}
func (r Recorder) Warn(args ...interface{}) {
	r.e.Warn(args...)
}
func (r Recorder) Error(args ...interface{}) {
	r.e.Error(args...)
}
func (r Recorder) Panic(args ...interface{}) {
	r.e.Panic(args...)
}
func (r Recorder) Fatal(args ...interface{}) {
	r.e.Fatal(args...)
}

func (r Recorder) Debugf(format string, args ...interface{}) {
	r.e.Debugf(format, args...)
}
func (r Recorder) Infof(format string, args ...interface{}) {
	r.e.Infof(format, args...)
}
func (r Recorder) Warnf(format string, args ...interface{}) {
	r.e.Warnf(format, args...)
}
func (r Recorder) Errorf(format string, args ...interface{}) {
	r.e.Errorf(format, args...)
}
func (r Recorder) Panicf(format string, args ...interface{}) {
	r.e.Panicf(format, args...)
}
func (r Recorder) Fatalf(format string, args ...interface{}) {
	r.e.Fatalf(format, args...)
}

func (r Recorder) Debugw(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Debug(msg)
}
func (r Recorder) Infow(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Info(msg)
}
func (r Recorder) Warnw(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Warn(msg)
}
func (r Recorder) Errorw(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Error(msg)
}
func (r Recorder) Panicw(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Panic(msg)
}
func (r Recorder) Fatalw(msg string, keyValuePairs ...interface{}) {
	r.e.WithFields(toFields(keyValuePairs)).Fatal(msg)
}

func (l Recorder) With(fields log.Fields) log.Recorder {
	return Recorder{e: l.e.WithFields(logrus.Fields(fields))}
}

type Logger struct {
	Recorder
	SyncFunc  func() error
	CloseFunc func() error
}

func NewLogger(l *logrus.Logger) Logger {
	return Logger{
		Recorder: Recorder{e: logrus.NewEntry(l)},
	}
}

func (l Logger) Sync() error {
	if l.SyncFunc == nil {
		return nil
	}
	return l.SyncFunc()
}
func (l Logger) Close() error {
	if l.CloseFunc == nil {
		return nil
	}
	return l.CloseFunc()
}
