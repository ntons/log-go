package log

var nop = nopLogger{}

type nopRecorder struct{}

func (nopRecorder) Debug(args ...interface{}) {}
func (nopRecorder) Info(args ...interface{})  {}
func (nopRecorder) Warn(args ...interface{})  {}
func (nopRecorder) Error(args ...interface{}) {}
func (nopRecorder) Panic(args ...interface{}) {}
func (nopRecorder) Fatal(args ...interface{}) {}

func (nopRecorder) Debugf(format string, args ...interface{}) {}
func (nopRecorder) Infof(format string, args ...interface{})  {}
func (nopRecorder) Warnf(format string, args ...interface{})  {}
func (nopRecorder) Errorf(format string, args ...interface{}) {}
func (nopRecorder) Panicf(format string, args ...interface{}) {}
func (nopRecorder) Fatalf(format string, args ...interface{}) {}

func (nopRecorder) Debugw(msg string, keyValuePairs ...interface{}) {}
func (nopRecorder) Infow(msg string, keyValuePairs ...interface{})  {}
func (nopRecorder) Warnw(msg string, keyValuePairs ...interface{})  {}
func (nopRecorder) Errorw(msg string, keyValuePairs ...interface{}) {}
func (nopRecorder) Panicw(msg string, keyValuePairs ...interface{}) {}
func (nopRecorder) Fatalw(msg string, keyValuePairs ...interface{}) {}

func (nopRecorder) With(Fields) Recorder { return nopRecorder{} }

type nopLogger struct{ nopRecorder }

func (nopLogger) Sync() error  { return nil }
func (nopLogger) Close() error { return nil }

// Check whether logger is nop
func IsNop(l Logger) bool {
	if l == nop {
		return true
	}
	switch l.(type) {
	case nopLogger, *nopLogger:
		return true
	default:
		return false
	}
}
