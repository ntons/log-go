package log

var Std LoggerLevelSetter

func SetStdLogger(logger LoggerLevelSetter) LoggerLevelSetter {
	oldStd := Std
	Std = logger
	return oldStd
}
func SetStdLevel(level Level) {
	Std.SetLevel(level)
}

func Debug(v ...interface{}) { Std.Debug(v...) }
func Info(v ...interface{})  { Std.Info(v...) }
func Warn(v ...interface{})  { Std.Warn(v...) }
func Error(v ...interface{}) { Std.Error(v...) }
func Fatal(v ...interface{}) { Std.Fatal(v...) }

func Debugf(format string, v ...interface{}) { Std.Debugf(format, v...) }
func Infof(format string, v ...interface{})  { Std.Infof(format, v...) }
func Warnf(format string, v ...interface{})  { Std.Warnf(format, v...) }
func Errorf(format string, v ...interface{}) { Std.Errorf(format, v...) }
func Fatalf(format string, v ...interface{}) { Std.Fatalf(format, v...) }

func Debugw(msg string, kvp ...interface{}) { Std.Debugw(msg, kvp...) }
func Infow(msg string, kvp ...interface{})  { Std.Infow(msg, kvp...) }
func Warnw(msg string, kvp ...interface{})  { Std.Warnw(msg, kvp...) }
func Errorw(msg string, kvp ...interface{}) { Std.Errorw(msg, kvp...) }
func Fatalw(msg string, kvp ...interface{}) { Std.Fatalw(msg, kvp...) }

func With(kvp ...interface{}) Logger { return addStack(Std.With(kvp...)) }

type wrap struct{ logger Logger }

func addStack(logger Logger) Logger {
	return wrap{logger}
}
func (w wrap) Debug(v ...interface{}) {
	w.logger.Debug(v...)
}
func (w wrap) Info(v ...interface{}) {
	w.logger.Info(v...)
}
func (w wrap) Warn(v ...interface{}) {
	w.logger.Warn(v...)
}
func (w wrap) Error(v ...interface{}) {
	w.logger.Error(v...)
}
func (w wrap) Panic(v ...interface{}) {
	w.logger.Panic(v...)
}
func (w wrap) Fatal(v ...interface{}) {
	w.logger.Fatal(v...)
}
func (w wrap) Debugf(format string, v ...interface{}) {
	w.logger.Debugf(format, v...)
}
func (w wrap) Infof(format string, v ...interface{}) {
	w.logger.Infof(format, v...)
}
func (w wrap) Warnf(format string, v ...interface{}) {
	w.logger.Warnf(format, v...)
}
func (w wrap) Errorf(format string, v ...interface{}) {
	w.logger.Errorf(format, v...)
}
func (w wrap) Panicf(format string, v ...interface{}) {
	w.logger.Panicf(format, v...)
}
func (w wrap) Fatalf(format string, v ...interface{}) {
	w.logger.Fatalf(format, v)
}
func (w wrap) Debugw(msg string, kvp ...interface{}) {
	w.logger.Debugw(msg, kvp...)
}
func (w wrap) Infow(msg string, kvp ...interface{}) {
	w.logger.Infow(msg, kvp...)
}
func (w wrap) Warnw(msg string, kvp ...interface{}) {
	w.logger.Warnw(msg, kvp...)
}
func (w wrap) Errorw(msg string, kvp ...interface{}) {
	w.logger.Errorw(msg, kvp...)
}
func (w wrap) Panicw(msg string, kvp ...interface{}) {
	w.logger.Panicw(msg, kvp...)
}
func (w wrap) Fatalw(msg string, kvp ...interface{}) {
	w.logger.Fatalw(msg, kvp...)
}
func (w wrap) With(kvp ...interface{}) Logger {
	return addStack(w.logger.With(kvp))
}
