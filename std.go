package log

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

var (
	ErrNotExist = fmt.Errorf("not exist")
)

var std unsafe.Pointer = unsafe.Pointer(newStd())

// atomic global variables
type stdAtom struct {
	D deeperLogger            // default
	L map[string]deeperLogger // loggers
}

func newStd() *stdAtom {
	return &stdAtom{D: deeperLogger{nop}, L: make(map[string]deeperLogger)}
}
func storeStd(new *stdAtom) {
	atomic.StorePointer(&std, unsafe.Pointer(new))
}
func swapStd(new *stdAtom) (old *stdAtom) {
	return (*stdAtom)(atomic.SwapPointer(&std, unsafe.Pointer(new)))
}
func loadStd() *stdAtom {
	return (*stdAtom)(atomic.LoadPointer(&std))
}

func ReplaceLoggers(m map[string]Logger, use string) {
	new := newStd()
	// wrap loggers for deeper callstack
	for name, l := range m {
		new.L[name] = deeperLogger{l}
		if name == use {
			new.D = deeperLogger{l}
		}
	}
	// replace std loggers
	old := swapStd(new).L
	// close old loggers after a delay
	time.AfterFunc(time.Minute, func() {
		for _, l := range old {
			l.Close()
		}
	})
}

// Check Logger existance
func HasLogger(name string) bool {
	_, ok := loadStd().L[name]
	return ok
}

// Get logger by name, return nop if not exist
func L(name string) Logger {
	if l, ok := loadStd().L[name]; ok {
		return l
	} else {
		return nop
	}
}

// default logger for further operating, return nop if not exist
func D() Logger {
	return loadStd().D
}

// use logger as default
func Use(name string) error {
	old := loadStd()
	for name_, l := range old.L {
		if name_ == name {
			storeStd(&stdAtom{D: l, L: old.L})
			return nil
		}
	}
	return ErrNotExist
}

// ONLY Recorder interface exported
func With(fields Fields) Recorder {
	return loadStd().D.With(fields)
}
func Debug(args ...interface{}) {
	loadStd().D.Logger.Debug(args...)
}
func Info(args ...interface{}) {
	loadStd().D.Logger.Info(args...)
}
func Warn(args ...interface{}) {
	loadStd().D.Logger.Warn(args...)
}
func Error(args ...interface{}) {
	loadStd().D.Logger.Error(args...)
}
func Fatal(args ...interface{}) {
	loadStd().D.Logger.Fatal(args...)
}
func Debugf(format string, args ...interface{}) {
	loadStd().D.Logger.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	loadStd().D.Logger.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	loadStd().D.Logger.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	loadStd().D.Logger.Errorf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	loadStd().D.Logger.Fatalf(format, args...)
}
func Debugw(msg string, fields Fields) {
	loadStd().D.Logger.Debugw(msg, fields)
}
func Infow(msg string, fields Fields) {
	loadStd().D.Logger.Infow(msg, fields)
}
func Warnw(msg string, fields Fields) {
	loadStd().D.Logger.Warnw(msg, fields)
}
func Errorw(msg string, fields Fields) {
	loadStd().D.Logger.Errorw(msg, fields)
}
func Fatalw(msg string, fields Fields) {
	loadStd().D.Logger.Fatalw(msg, fields)
}
