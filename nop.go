package log

type NopRecorder struct {
}

func (n *NopRecorder) With(fields Fields) Recorder            { return n }
func (NopRecorder) Debug(args ...interface{})                 {}
func (NopRecorder) Info(args ...interface{})                  {}
func (NopRecorder) Warn(args ...interface{})                  {}
func (NopRecorder) Error(args ...interface{})                 {}
func (NopRecorder) Fatal(args ...interface{})                 {}
func (NopRecorder) Debugf(format string, args ...interface{}) {}
func (NopRecorder) Infof(format string, args ...interface{})  {}
func (NopRecorder) Warnf(format string, args ...interface{})  {}
func (NopRecorder) Errorf(format string, args ...interface{}) {}
func (NopRecorder) Fatalf(format string, args ...interface{}) {}
func (NopRecorder) Debugw(msg string, fields Fields)          {}
func (NopRecorder) Infow(msg string, fields Fields)           {}
func (NopRecorder) Warnw(msg string, fields Fields)           {}
func (NopRecorder) Errorw(msg string, fields Fields)          {}
func (NopRecorder) Fatalw(msg string, fields Fields)          {}

type NopLogger struct {
	NopRecorder
}

func (NopLogger) Write(p []byte) (int, error)   { return len(p), nil }
func (NopLogger) Sync() error                   { return nil }
func (NopLogger) Close() error                  { return nil }
func (NopLogger) SetLevel(lev Level)            {}
func (NopLogger) IsLevelEnabled(lev Level) bool { return false }
