package log

import (
	"sync/atomic"
	"unsafe"
)

// wrap logger to pointer, the pointer could be swapped atomically
type wrap struct{ Logger }

var std = unsafe.Pointer(&wrap{StdLogger{}})

func SetLogger(l Logger) Logger {
	return (*wrap)(atomic.SwapPointer(&std, unsafe.Pointer(&wrap{l})))
}
func getStd() Logger {
	return (*wrap)(atomic.LoadPointer(&std))
}

func Debug(v ...interface{}) { getStd().Debug(v...) }
func Info(v ...interface{})  { getStd().Info(v...) }
func Warn(v ...interface{})  { getStd().Warn(v...) }
func Error(v ...interface{}) { getStd().Error(v...) }
func Fatal(v ...interface{}) { getStd().Fatal(v...) }

func Debugf(format string, v ...interface{}) { getStd().Debugf(format, v...) }
func Infof(format string, v ...interface{})  { getStd().Infof(format, v...) }
func Warnf(format string, v ...interface{})  { getStd().Warnf(format, v...) }
func Errorf(format string, v ...interface{}) { getStd().Errorf(format, v...) }
func Fatalf(format string, v ...interface{}) { getStd().Fatalf(format, v...) }

func Debugw(msg string, kvp ...interface{}) { getStd().Debugw(msg, kvp...) }
func Infow(msg string, kvp ...interface{})  { getStd().Infow(msg, kvp...) }
func Warnw(msg string, kvp ...interface{})  { getStd().Warnw(msg, kvp...) }
func Errorw(msg string, kvp ...interface{}) { getStd().Errorw(msg, kvp...) }
func Fatalw(msg string, kvp ...interface{}) { getStd().Fatalw(msg, kvp...) }

func With(kvp ...interface{}) Logger { return stack{getStd().With(kvp...)} }
