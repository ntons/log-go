// Custom Logger MUST skip 1 more caller stack for std wrapper

package log

import (
	"io"
)

type Fields map[string]interface{}

type Syncer interface {
	Sync() error
}

type Recorder interface {
	// msg = fmt.Sprint(args...)
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Fatal(args ...interface{})

	// msg = fmt.Sprintf(format, args...)
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	// structured
	Debugw(msg string, fields Fields)
	Infow(msg string, fields Fields)
	Warnw(msg string, fields Fields)
	Errorw(msg string, fields Fields)
	Panicw(msg string, fields Fields)
	Fatalw(msg string, fields Fields)

	// preset fields
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
