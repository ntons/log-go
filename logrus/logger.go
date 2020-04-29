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
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Debug(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Debug(args...)
	}
}
func (e Recorder) Info(args ...interface{}) {
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Info(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Info(args...)
	}
}
func (e Recorder) Warn(args ...interface{}) {
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Warn(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Warn(args...)
	}
}
func (e Recorder) Error(args ...interface{}) {
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Error(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Error(args...)
	}
}
func (e Recorder) Panic(args ...interface{}) {
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Panic(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Panic(args...)
	}
}
func (e Recorder) Fatal(args ...interface{}) {
	fields, args := log.ExtractFields(args)
	if len(fields) == 0 {
		e.e.Fatal(args...)
	} else {
		e.e.WithFields(logrus.Fields(fields)).Fatal(args...)
	}
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
