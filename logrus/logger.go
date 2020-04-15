package logrus

import (
	"github.com/ntons/log-go"
	"github.com/sirupsen/logrus"
)

func fitFields(fields log.Fields) []interface{} {
	arr := make([]interface{}, 0, len(fields)*2)
	for key, val := range fields {
		arr = append(arr, key, val)
	}
	return arr
}

func fitLevel(lev log.Level) logrus.Level {
	switch lev {
	case log.DebugLevel:
		return logrus.DebugLevel
	case log.InfoLevel:
		return logrus.InfoLevel
	case log.WarnLevel:
		return logrus.WarnLevel
	case log.ErrorLevel:
		return logrus.ErrorLevel
	case log.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.PanicLevel
	}
}

type Recorder struct {
	e *logrus.Entry
}

func (l Recorder) With(fields log.Fields) log.Recorder {
	return Recorder{
		e: l.e.WithFields(logrus.Fields(fields)),
	}
}

func (e Recorder) Debug(args ...interface{}) {
	e.e.Debug(args...)
}
func (e Recorder) Info(args ...interface{}) {
	e.e.Info(args...)
}
func (e Recorder) Warn(args ...interface{}) {
	e.e.Warn(args...)
}
func (e Recorder) Error(args ...interface{}) {
	e.e.Error(args...)
}
func (e Recorder) Panic(args ...interface{}) {
	e.e.Panic(args...)
}
func (e Recorder) Fatal(args ...interface{}) {
	e.e.Fatal(args...)
}

func (e Recorder) Debugf(format string, args ...interface{}) {
	e.e.Debugf(format, args...)
}
func (e Recorder) Infof(format string, args ...interface{}) {
	e.e.Infof(format, args...)
}
func (e Recorder) Warnf(format string, args ...interface{}) {
	e.e.Warnf(format, args...)
}
func (e Recorder) Errorf(format string, args ...interface{}) {
	e.e.Errorf(format, args...)
}
func (e Recorder) Panicf(format string, args ...interface{}) {
	e.e.Panicf(format, args...)
}
func (e Recorder) Fatalf(format string, args ...interface{}) {
	e.e.Fatalf(format, args...)
}

func (e Recorder) Debugw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Debug(msg)
}
func (e Recorder) Infow(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Info(msg)
}
func (e Recorder) Warnw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Warn(msg)
}
func (e Recorder) Errorw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Error(msg)
}
func (e Recorder) Panicw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Panic(msg)
}
func (e Recorder) Fatalw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Fatal(msg)
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
