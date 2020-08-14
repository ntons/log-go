// Custom Logger MUST skip 1 more caller stack for std wrapper

package log

type Fields map[string]interface{}
type M = Fields

type Syncer interface {
	Sync() error
}

type Recorder interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debugw(msg string, keyValuePairs ...interface{})
	Infow(msg string, keyValuePairs ...interface{})
	Warnw(msg string, keyValuePairs ...interface{})
	Errorw(msg string, keyValuePairs ...interface{})
	Panicw(msg string, keyValuePairs ...interface{})
	Fatalw(msg string, keyValuePairs ...interface{})

	With(fields Fields) Recorder
	Withw(keyValuePairs ...interface{}) Recorder
}

type Logger interface {
	Syncer
	Recorder
}
