package log

var nop = nopLogger{}

type nopRecorder struct{}

func (n nopRecorder) With(fields Fields) Recorder             { return n }
func (nopRecorder) Debug(args ...interface{})                 {}
func (nopRecorder) Info(args ...interface{})                  {}
func (nopRecorder) Warn(args ...interface{})                  {}
func (nopRecorder) Error(args ...interface{})                 {}
func (nopRecorder) Fatal(args ...interface{})                 {}
func (nopRecorder) Debugf(format string, args ...interface{}) {}
func (nopRecorder) Infof(format string, args ...interface{})  {}
func (nopRecorder) Warnf(format string, args ...interface{})  {}
func (nopRecorder) Errorf(format string, args ...interface{}) {}
func (nopRecorder) Fatalf(format string, args ...interface{}) {}
func (nopRecorder) Debugw(msg string, fields Fields)          {}
func (nopRecorder) Infow(msg string, fields Fields)           {}
func (nopRecorder) Warnw(msg string, fields Fields)           {}
func (nopRecorder) Errorw(msg string, fields Fields)          {}
func (nopRecorder) Fatalw(msg string, fields Fields)          {}

type nopLogger struct{ nopRecorder }

func (nopLogger) Write(p []byte) (int, error)   { return len(p), nil }
func (nopLogger) Sync() error                   { return nil }
func (nopLogger) Close() error                  { return nil }
func (nopLogger) SetLevel(lev Level)            {}
func (nopLogger) IsLevelEnabled(lev Level) bool { return false }

// Check whether logger is nop
func IsNop(l Logger) bool {
	if l == nop {
		return true
	}
	switch x := l.(type) {
	case nopLogger:
		return true
	case *nopLogger:
		return true
	case deeperLogger:
		return IsNop(x.Logger)
	case *deeperLogger:
		return IsNop(x.Logger)
	default:
		return false
	}
}
