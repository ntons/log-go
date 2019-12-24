package log

import (
	"fmt"
	"io"
	"strings"
	"sync"
)

type Fields map[string]interface{}

type Syncer interface {
	Sync() error
}

type SyncCloser interface {
	Syncer
	io.Closer
}

type Recorder interface {
	// preset context
	With(fields Fields) Recorder

	// msg = fmt.Sprint(args...)
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	// msg = fmt.Sprintf(format, args...)
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	// structured
	Debugw(msg string, fields Fields)
	Infow(msg string, fields Fields)
	Warnw(msg string, fields Fields)
	Errorw(msg string, fields Fields)
	Fatalw(msg string, fields Fields)
}

type Logger interface {
	Recorder
	SyncCloser

	SetLevel(lev Level)
	IsLevelEnabled(lev Level) bool
}

// JSON unmarshallable logger builder
type Builder interface {
	Build() (Logger, error)
}

var (
	// DONOT usage these variables directly
	_factoryMtx sync.Mutex
	_factoryMap = make(map[string]func() Builder)
)

func RegisterLogger(name string, factory func() Builder) error {
	name = strings.ToLower(name) // ignore case
	_factoryMtx.Lock()
	defer _factoryMtx.Unlock()
	if _, ok := _factoryMap[name]; ok {
		return fmt.Errorf("logger already registered for name %q", name)
	}
	_factoryMap[name] = factory
	return nil
}

func newLoggerBuilder(name string) (Builder, error) {
	name = strings.ToLower(name)
	_factoryMtx.Lock()
	defer _factoryMtx.Unlock()
	factory, ok := _factoryMap[name]
	if !ok {
		return nil, fmt.Errorf("no logger registered for name %q", name)
	}
	return factory(), nil
}
