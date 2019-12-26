package log

// recorder deeperLogger wrapper
type deeperRecorder struct {
	Recorder
}

func (d deeperRecorder) With(fields Fields) Recorder {
	return deeperRecorder{d.Recorder.With(fields)}
}
func (d deeperRecorder) Debug(args ...interface{}) {
	d.Recorder.Debug(args...)
}
func (d deeperRecorder) Info(args ...interface{}) {
	d.Recorder.Info(args...)
}
func (d deeperRecorder) Warn(args ...interface{}) {
	d.Recorder.Warn(args...)
}
func (d deeperRecorder) Error(args ...interface{}) {
	d.Recorder.Error(args...)
}
func (d deeperRecorder) Fatal(args ...interface{}) {
	d.Recorder.Fatal(args...)
}
func (d deeperRecorder) Debugf(format string, args ...interface{}) {
	d.Recorder.Debugf(format, args...)
}
func (d deeperRecorder) Infof(format string, args ...interface{}) {
	d.Recorder.Infof(format, args...)
}
func (d deeperRecorder) Warnf(format string, args ...interface{}) {
	d.Warnf(format, args...)
}
func (d deeperRecorder) Errorf(format string, args ...interface{}) {
	d.Errorf(format, args...)
}
func (d deeperRecorder) Fatalf(format string, args ...interface{}) {
	d.Recorder.Fatalf(format, args...)
}
func (d deeperRecorder) Debugw(msg string, fields Fields) {
	d.Recorder.Debugw(msg, fields)
}
func (d deeperRecorder) Infow(msg string, fields Fields) {
	d.Recorder.Infow(msg, fields)
}
func (d deeperRecorder) Warnw(msg string, fields Fields) {
	d.Recorder.Warnw(msg, fields)
}
func (d deeperRecorder) Errorw(msg string, fields Fields) {
	d.Recorder.Errorw(msg, fields)
}
func (d deeperRecorder) Fatalw(msg string, fields Fields) {
	d.Recorder.Fatalw(msg, fields)
}

// logger deeperLogger wrapper
type deeperLogger struct {
	Logger
}

func (d deeperLogger) With(fields Fields) Recorder {
	return &deeperRecorder{d.Logger.With(fields)}
}
func (d deeperLogger) Debug(args ...interface{}) {
	d.Logger.Debug(args...)
}
func (d deeperLogger) Info(args ...interface{}) {
	d.Logger.Info(args...)
}
func (d deeperLogger) Warn(args ...interface{}) {
	d.Logger.Warn(args...)
}
func (d deeperLogger) Error(args ...interface{}) {
	d.Logger.Error(args...)
}
func (d deeperLogger) Fatal(args ...interface{}) {
	d.Logger.Fatal(args...)
}
func (d deeperLogger) Debugf(format string, args ...interface{}) {
	d.Logger.Debugf(format, args...)
}
func (d deeperLogger) Infof(format string, args ...interface{}) {
	d.Logger.Infof(format, args...)
}
func (d deeperLogger) Warnf(format string, args ...interface{}) {
	d.Warnf(format, args...)
}
func (d deeperLogger) Errorf(format string, args ...interface{}) {
	d.Errorf(format, args...)
}
func (d deeperLogger) Fatalf(format string, args ...interface{}) {
	d.Logger.Fatalf(format, args...)
}
func (d deeperLogger) Debugw(msg string, fields Fields) {
	d.Logger.Debugw(msg, fields)
}
func (d deeperLogger) Infow(msg string, fields Fields) {
	d.Logger.Infow(msg, fields)
}
func (d deeperLogger) Warnw(msg string, fields Fields) {
	d.Logger.Warnw(msg, fields)
}
func (d deeperLogger) Errorw(msg string, fields Fields) {
	d.Logger.Errorw(msg, fields)
}
func (d deeperLogger) Fatalw(msg string, fields Fields) {
	d.Logger.Fatalw(msg, fields)
}
