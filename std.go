package log

import (
	"sync/atomic"
	"unsafe"
)

type globals struct {
	Std     Logger
	Loggers map[string]Logger
}

var g unsafe.Pointer = unsafe.Pointer(&globals{
	Std: nil, Loggers: make(map[string]Logger),
})

func replaceGlobalLoggers(loggers map[string]Logger) {
	atomic.StorePointer(&g, unsafe.Pointer(&globals{
		Std:     loggers["std"],
		Loggers: loggers,
	}))
}

// Get logger by name, return nil if not exist
func L(name string) Logger {
	return (*globals)(atomic.LoadPointer(&g)).Loggers[name]
}

func Foreach(f func(string, Logger) error) (err error) {
	for name, logger := range (*globals)(atomic.LoadPointer(&g)).Loggers {
		if err = f(name, logger); err != nil {
			return
		}
	}
	return
}

// sync ALL logger and return latest error
func SyncAll() (err error) {
	Foreach(func(_ string, logger Logger) error {
		if syncErr := logger.Sync(); syncErr != nil {
			err = syncErr
		}
		return nil
	})
	return
}

// close ALL logger and return latest error
func CloseAll() (err error) {
	Foreach(func(_ string, logger Logger) error {
		if closeErr := logger.Close(); closeErr != nil {
			err = closeErr
		}
		return nil
	})
	return
}

// standard logger for further operating
func Std() Logger {
	return (*globals)(atomic.LoadPointer(&g)).Std
}

// ONLY Recorder interface exported
func With(fields Fields) Recorder {
	return Std().With(fields)
}

func Debug(args ...interface{}) {
	Std().Debug(args...)
}
func Info(args ...interface{}) {
	Std().Info(args...)
}
func Warn(args ...interface{}) {
	Std().Warn(args...)
}
func Error(args ...interface{}) {
	Std().Error(args...)
}
func Fatal(args ...interface{}) {
	Std().Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	Std().Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	Std().Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	Std().Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	Std().Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	Std().Fatalf(format, args...)
}

func Debugw(msg string, fields Fields) {
	Std().Debugw(msg, fields)
}
func Infow(msg string, fields Fields) {
	Std().Infow(msg, fields)
}
func Warnw(msg string, fields Fields) {
	Std().Warnw(msg, fields)
}
func Errorw(msg string, fields Fields) {
	Std().Errorw(msg, fields)
}
func Fatalw(msg string, fields Fields) {
	Std().Fatalw(msg, fields)
}
