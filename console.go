package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type consoleLoggerOptions struct {
	level Level
	color bool
}

type ConsoleLoggerOption interface {
	apply(o *consoleLoggerOptions)
}

type funcConsoleLoggerOption struct {
	f func(o *consoleLoggerOptions)
}

func (f funcConsoleLoggerOption) apply(o *consoleLoggerOptions) {
	f.f(o)
}

func ConsoleLoggerWithLevel(level Level) ConsoleLoggerOption {
	return funcConsoleLoggerOption{f: func(o *consoleLoggerOptions) {
		o.level = level
	}}
}

func ConsoleLoggerWithColor() ConsoleLoggerOption {
	return funcConsoleLoggerOption{f: func(o *consoleLoggerOptions) {
		o.color = true
	}}
}

var consolemu sync.Mutex

type ConsoleLogger struct {
	*consoleLoggerOptions
	fields Fields
}

func NewConsoleLogger(opts ...ConsoleLoggerOption) *ConsoleLogger {
	o := &consoleLoggerOptions{}
	for _, opt := range opts {
		opt.apply(o)
	}
	return &ConsoleLogger{consoleLoggerOptions: o}
}

func (l *ConsoleLogger) write(fields Fields) {
	_, file, line, _ := runtime.Caller(3)
	for key, val := range l.fields {
		fields[key] = val
	}
	lev := fields["lev"]
	delete(fields, "lev")
	msg := fields["msg"]
	delete(fields, "msg")

	consolemu.Lock()
	defer consolemu.Unlock()

	if l.color {
		switch lev {
		case "I":
			fmt.Fprint(os.Stdout, "\033[32m")
			defer fmt.Fprint(os.Stdout, "\033[0m")
		case "W":
			fmt.Fprint(os.Stdout, "\033[33m")
			defer fmt.Fprint(os.Stdout, "\033[0m")
		case "E", "P", "F":
			fmt.Fprint(os.Stdout, "\033[31m")
			defer fmt.Fprint(os.Stdout, "\033[0m")
		}
	}
	fmt.Fprintf(os.Stdout, "[%s]", time.Now().Format("2006-01-02 15:04:05.000"))
	fmt.Fprintf(os.Stdout, "[%s]", lev)
	fmt.Fprintf(os.Stdout, "[%s:%d]", filepath.Base(file), line)
	for key, val := range fields {
		fmt.Fprintf(os.Stdout, "[%s:%v]", key, val)
	}
	fmt.Fprintf(os.Stdout, " %s\n", msg)
}

func (l ConsoleLogger) Close() error {
	return l.Sync()
}
func (l ConsoleLogger) Sync() error {
	return os.Stdout.Sync()
}
func (l ConsoleLogger) Debug(args ...interface{}) {
	if l.level <= DebugLevel {
		l.write(Fields{"lev": "D", "msg": fmt.Sprint(args...)})
	}
}
func (l ConsoleLogger) Info(args ...interface{}) {
	if l.level <= InfoLevel {
		l.write(Fields{"lev": "I", "msg": fmt.Sprint(args...)})
	}
}
func (l ConsoleLogger) Warn(args ...interface{}) {
	if l.level <= WarnLevel {
		l.write(Fields{"lev": "W", "msg": fmt.Sprint(args...)})
	}
}
func (l ConsoleLogger) Error(args ...interface{}) {
	if l.level <= ErrorLevel {
		l.write(Fields{"lev": "E", "msg": fmt.Sprint(args...)})
	}
}
func (l ConsoleLogger) Panic(args ...interface{}) {
	if l.level <= PanicLevel {
		l.write(Fields{"lev": "P", "msg": fmt.Sprint(args...)})
		panic(fmt.Sprint(args...))
	}
}
func (l ConsoleLogger) Fatal(args ...interface{}) {
	if l.level <= FatalLevel {
		l.write(Fields{"lev": "F", "msg": fmt.Sprint(args...)})
		os.Exit(1)
	}
}
func (l ConsoleLogger) Debugf(format string, args ...interface{}) {
	if l.level <= DebugLevel {
		l.write(Fields{"lev": "D", "msg": fmt.Sprintf(format, args...)})
	}
}
func (l ConsoleLogger) Infof(format string, args ...interface{}) {
	if l.level <= InfoLevel {
		l.write(Fields{"lev": "I", "msg": fmt.Sprintf(format, args...)})
	}
}
func (l ConsoleLogger) Warnf(format string, args ...interface{}) {
	if l.level <= WarnLevel {
		l.write(Fields{"lev": "W", "msg": fmt.Sprintf(format, args...)})
	}
}
func (l ConsoleLogger) Errorf(format string, args ...interface{}) {
	if l.level <= ErrorLevel {
		l.write(Fields{"lev": "E", "msg": fmt.Sprintf(format, args...)})
	}
}
func (l ConsoleLogger) Panicf(format string, args ...interface{}) {
	if l.level <= PanicLevel {
		l.write(Fields{"lev": "P", "msg": fmt.Sprintf(format, args...)})
		panic(fmt.Sprintf(format, args...))
	}
}
func (l ConsoleLogger) Fatalf(format string, args ...interface{}) {
	if l.level <= FatalLevel {
		l.write(Fields{"lev": "F", "msg": fmt.Sprintf(format, args...)})
		os.Exit(1)
	}
}

func mergeKeyValuePairsToFields(
	fields Fields, keyValuePairs []interface{}) Fields {
	for i := 1; i < len(keyValuePairs); i += 2 {
		key, ok := keyValuePairs[i-1].(string)
		if !ok {
			key = fmt.Sprintf("%v", keyValuePairs[i-1])
		}
		fields[key] = keyValuePairs[i]
	}
	return fields
}
func (l ConsoleLogger) Debugw(msg string, keyValuePairs ...interface{}) {
	if l.level <= DebugLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "D", "msg": msg},
			keyValuePairs))
	}
}
func (l ConsoleLogger) Infow(msg string, keyValuePairs ...interface{}) {
	if l.level <= InfoLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "I", "msg": msg},
			keyValuePairs))
	}
}
func (l ConsoleLogger) Warnw(msg string, keyValuePairs ...interface{}) {
	if l.level <= WarnLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "W", "msg": msg},
			keyValuePairs))
	}
}
func (l ConsoleLogger) Errorw(msg string, keyValuePairs ...interface{}) {
	if l.level <= ErrorLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "E", "msg": msg},
			keyValuePairs))
	}
}
func (l ConsoleLogger) Panicw(msg string, keyValuePairs ...interface{}) {
	if l.level <= PanicLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "P", "msg": msg},
			keyValuePairs))
		panic(msg)
	}
}
func (l ConsoleLogger) Fatalw(msg string, keyValuePairs ...interface{}) {
	if l.level <= FatalLevel {
		l.write(mergeKeyValuePairsToFields(
			Fields{"lev": "F", "msg": msg},
			keyValuePairs))
		os.Exit(1)
	}
}

func (l ConsoleLogger) With(fields Fields) Recorder {
	for key, val := range l.fields {
		fields[key] = val
	}
	clone := l
	clone.fields = fields
	return clone
}
