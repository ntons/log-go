// Custom Logger MUST skip 1 more caller stack for std wrapper

package log

import (
	"io"
)

type Fields map[string]interface{}
type F = Fields // for short

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
}

type Logger interface {
	// logger could be closed after replacing
	io.Closer
	// flush all data to destination
	Syncer
	// logger is also a recorder
	Recorder
}
