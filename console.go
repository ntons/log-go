package log

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type ConsoleLogger struct {
	fields Fields
	color  bool
}

func NewConsoleLogger(color bool) ConsoleLogger {
	return ConsoleLogger{color: color}
}

func (l ConsoleLogger) write(fields Fields) {
	_, file, line, _ := runtime.Caller(3)
	for key, val := range l.fields {
		fields[key] = val
	}
	lev := fields["lev"]
	delete(fields, "lev")
	msg := fields["msg"]
	delete(fields, "msg")

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
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "D",
		"msg": fmt.Sprint(args...),
	}))
}
func (l ConsoleLogger) Info(args ...interface{}) {
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "I",
		"msg": fmt.Sprint(args...),
	}))
}
func (l ConsoleLogger) Warn(args ...interface{}) {
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "W",
		"msg": fmt.Sprint(args...),
	}))
}
func (l ConsoleLogger) Error(args ...interface{}) {
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "E",
		"msg": fmt.Sprint(args...),
	}))
}
func (l ConsoleLogger) Panic(args ...interface{}) {
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "P",
		"msg": fmt.Sprint(args...),
	}))
	panic(fmt.Sprint(args...))
}
func (l ConsoleLogger) Fatal(args ...interface{}) {
	fields, args := ExtractFields(args)
	l.write(MergeFields(fields, Fields{
		"lev": "F",
		"msg": fmt.Sprint(args...),
	}))
	os.Exit(1)
}
func (l ConsoleLogger) Debugf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "D",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l ConsoleLogger) Infof(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "I",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l ConsoleLogger) Warnf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "W",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l ConsoleLogger) Errorf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "E",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l ConsoleLogger) Panicf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "P",
		"msg": fmt.Sprintf(format, args...),
	})
	panic(fmt.Sprintf(format, args...))
}
func (l ConsoleLogger) Fatalf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "F",
		"msg": fmt.Sprintf(format, args...),
	})
	os.Exit(1)
}

func (l ConsoleLogger) With(fields Fields) Recorder {
	for key, val := range l.fields {
		fields[key] = val
	}
	return ConsoleLogger{fields: fields}
}
