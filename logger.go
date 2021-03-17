// Custom Logger MUST skip 1 more caller stack for std wrapper

package log

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warn(v ...interface{})
	Error(v ...interface{})
	Panic(v ...interface{})
	Fatal(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})

	Debugw(msg string, kvp ...interface{})
	Infow(msg string, kvp ...interface{})
	Warnw(msg string, kvp ...interface{})
	Errorw(msg string, kvp ...interface{})
	Panicw(msg string, kvp ...interface{})
	Fatalw(msg string, kvp ...interface{})

	With(kvp ...interface{}) Logger
}
