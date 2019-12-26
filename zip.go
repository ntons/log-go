package log

import (
	"fmt"
)

// After replacing, the old std will be closed, all Recorder allocated
// from old std fail. so we need reallocate them from the new std, and
// the Recorder object which user holds could still be used silently.
// The With operation makes a tree, uddate could backtrace from one
// note to root, check and update reference to underlying recorder.
// The updated node can be reused by update from other branches.

type Zipper interface {
	with(fields Fields) Recorder
}

// recorder deeperLogger wrapper
type Zip struct {
	// associated logger with recorder
	l Logger
	r Recorder
	p Zipper
	f Fields
}

func newZip(l Logger, fields Fields) Recorder {
	return &Zip{
		l: l,
		p: nil,
		f: fields,
		r: l.With(fields),
	}
}
func (z *Zip) check() {
	if z.l != std {
		fmt.Printf("zip updated %v\n", z.f)
		z.l = std
		if z.p != nil {
			z.r = z.p.with(z.f)
		} else {
			z.r = z.l.With(z.f)
		}
	}
}
func (z *Zip) with(fields Fields) Recorder {
	z.check()
	return &Zip{l: z.l, p: z, f: fields, r: z.r.With(fields)}
}

func (z *Zip) With(fields Fields) Recorder {
	mu.RLock()
	defer mu.RUnlock()
	return z.with(fields)
}
func (z *Zip) Debug(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Debug(args...)
}
func (z *Zip) Info(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Info(args...)
}
func (z *Zip) Warn(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Warn(args...)
}
func (z *Zip) Error(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Error(args...)
}
func (z *Zip) Panic(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Panic(args...)
}
func (z *Zip) Fatal(args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Fatal(args...)
}
func (z *Zip) Debugf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Debugf(format, args...)
}
func (z *Zip) Infof(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Infof(format, args...)
}
func (z *Zip) Warnf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Warnf(format, args...)
}
func (z *Zip) Errorf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Errorf(format, args...)
}
func (z *Zip) Panicf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Panicf(format, args...)
}
func (z *Zip) Fatalf(format string, args ...interface{}) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Fatalf(format, args...)
}
func (z *Zip) Debugw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Debugw(msg, fields)
}
func (z *Zip) Infow(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Infow(msg, fields)
}
func (z *Zip) Warnw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Warnw(msg, fields)
}
func (z *Zip) Errorw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Errorw(msg, fields)
}
func (z *Zip) Panicw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Panicw(msg, fields)
}
func (z *Zip) Fatalw(msg string, fields Fields) {
	mu.RLock()
	defer mu.RUnlock()
	z.check()
	z.r.Fatalw(msg, fields)
}
