package log

// add call-stack to logger for global With
type stack struct{ l Logger }

func (x stack) Debug(v ...interface{}) {
	x.l.Debug(v...)
}
func (x stack) Info(v ...interface{}) {
	x.l.Info(v...)
}
func (x stack) Warn(v ...interface{}) {
	x.l.Warn(v...)
}
func (x stack) Error(v ...interface{}) {
	x.l.Error(v...)
}
func (x stack) Panic(v ...interface{}) {
	x.l.Panic(v...)
}
func (x stack) Fatal(v ...interface{}) {
	x.l.Fatal(v...)
}

func (x stack) Debugf(format string, v ...interface{}) {
	x.l.Debugf(format, v...)
}
func (x stack) Infof(format string, v ...interface{}) {
	x.l.Infof(format, v...)
}
func (x stack) Warnf(format string, v ...interface{}) {
	x.l.Warnf(format, v...)
}
func (x stack) Errorf(format string, v ...interface{}) {
	x.l.Errorf(format, v...)
}
func (x stack) Panicf(format string, v ...interface{}) {
	x.l.Panicf(format, v...)
}
func (x stack) Fatalf(format string, v ...interface{}) {
	x.l.Fatalf(format, v)
}

func (x stack) Debugw(msg string, kvp ...interface{}) {
	x.l.Debugw(msg, kvp...)
}
func (x stack) Infow(msg string, kvp ...interface{}) {
	x.l.Infow(msg, kvp...)
}
func (x stack) Warnw(msg string, kvp ...interface{}) {
	x.l.Warnw(msg, kvp...)
}
func (x stack) Errorw(msg string, kvp ...interface{}) {
	x.l.Errorw(msg, kvp...)
}
func (x stack) Panicw(msg string, kvp ...interface{}) {
	x.l.Panicw(msg, kvp...)
}
func (x stack) Fatalw(msg string, kvp ...interface{}) {
	x.l.Fatalw(msg, kvp...)
}

func (x stack) With(kvp ...interface{}) Logger {
	return &stack{l: x.l.With(kvp...)}
}
