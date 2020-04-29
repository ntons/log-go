package log

import (
	"sync"
)

// convert "std" to pointer, so that can be comparable in zip
type box struct{ Logger }

var (
	std *box = &box{NewConsoleLogger(true)}
	// prevent logging while replacing
	mu sync.RWMutex
)

func ReplaceStd(new Logger) (old Logger) {
	mu.Lock() // the only place require write lock
	defer mu.Unlock()
	old, std = std.Logger, &box{new}
	return
}

func Close() error {
	return ReplaceStd(nop).Close()
}
func Sync() error {
	mu.RLock()
	defer mu.RUnlock()
	return std.Sync()
}
func Debug(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Debug(args...)
}
func Info(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Info(args...)
}
func Warn(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Warn(args...)
}
func Error(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Error(args...)
}
func Fatal(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Fatal(args...)
}

func Debugf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	std.Fatalf(format, args...)
}

func With(fields Fields) Recorder {
	mu.RLock()
	defer mu.RUnlock()
	return newWrap(std, fields)
}
