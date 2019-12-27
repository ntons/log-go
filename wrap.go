package log

import (
	"fmt"
)

// After replacing, the old std will be closed, all Recorder allocated
// from old std fail. so we need reallocate them from the new std, and
// the Recorder object which user holds could still be used correctly.
// The With operation makes a tree, update could backtrace from one
// note to root, check and update reference to underlying recorder.
// The updated node can be reused by update from other branches.

type zipper interface {
	with(fields Fields) Recorder
}

// recorder deeperLogger wrapper
type wrap struct {
	// associated logger with recorder
	l Logger
	r Recorder
	p zipper
	f Fields
}

func newWrap(l Logger, fields Fields) Recorder {
	return &wrap{
		l: l,
		p: nil,
		f: fields,
		r: l.With(fields),
	}
}
func (w *wrap) check() {
	if w.l != std {
		fmt.Printf("zip updated %v\n", w.f)
		w.l = std
		if w.p != nil {
			w.r = w.p.with(w.f)
		} else {
			w.r = w.l.With(w.f)
		}
	}
}
func (w *wrap) with(fields Fields) Recorder {
	w.check()
	return &wrap{l: w.l, p: w, f: fields, r: w.r.With(fields)}
}

func (w *wrap) With(fields Fields) Recorder {
	mu.RLock()
	defer mu.RUnlock()
	return w.with(fields)
}
func (w *wrap) Debug(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Debug(args...)
}
func (w *wrap) Info(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Info(args...)
}
func (w *wrap) Warn(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Warn(args...)
}
func (w *wrap) Error(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Error(args...)
}
func (w *wrap) Panic(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Panic(args...)
}
func (w *wrap) Fatal(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Fatal(args...)
}
func (w *wrap) Debugf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Debugf(format, args...)
}
func (w *wrap) Infof(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Infof(format, args...)
}
func (w *wrap) Warnf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Warnf(format, args...)
}
func (w *wrap) Errorf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Errorf(format, args...)
}
func (w *wrap) Panicf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Panicf(format, args...)
}
func (w *wrap) Fatalf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Fatalf(format, args...)
}
func (w *wrap) Debugw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Debugw(msg, fields)
}
func (w *wrap) Infow(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Infow(msg, fields)
}
func (w *wrap) Warnw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Warnw(msg, fields)
}
func (w *wrap) Errorw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Errorw(msg, fields)
}
func (w *wrap) Panicw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Panicw(msg, fields)
}
func (w *wrap) Fatalw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	w.check()
	w.r.Fatalw(msg, fields)
}
