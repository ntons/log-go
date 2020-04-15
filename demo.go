package log

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type DemoLogger struct {
	Out    io.Writer
	Name   string
	Fields Fields
}

func NewDemoLogger(out io.Writer, name string) DemoLogger {
	return DemoLogger{Out: out, Name: name}
}

func (l DemoLogger) write(fields Fields) {
	for key, val := range l.Fields {
		fields[key] = val
	}
	lev := fields["lev"]
	delete(fields, "lev")
	b, err := json.Marshal(fields)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(
		l.Out, "%s %s[%s] %s\n",
		time.Now().Format("2006-01-02 15:04:05.000"),
		l.Name, lev, b)
}

func (l DemoLogger) Close() error {
	fmt.Printf("DemoLogger %q closed\n", l.Name)
	return nil
}
func (l DemoLogger) Sync() error {
	return nil
}
func (l DemoLogger) Debug(args ...interface{}) {
	l.write(Fields{
		"lev": "D",
		"msg": fmt.Sprint(args...),
	})
}
func (l DemoLogger) Info(args ...interface{}) {
	l.write(Fields{
		"lev": "I",
		"msg": fmt.Sprint(args...),
	})
}
func (l DemoLogger) Warn(args ...interface{}) {
	l.write(Fields{
		"lev": "W",
		"msg": fmt.Sprint(args...),
	})
}
func (l DemoLogger) Error(args ...interface{}) {
	l.write(Fields{
		"lev": "E",
		"msg": fmt.Sprint(args...),
	})
}
func (l DemoLogger) Panic(args ...interface{}) {
	l.write(Fields{
		"lev": "P",
		"msg": fmt.Sprint(args...),
	})
	panic(fmt.Sprint(args...))
}
func (l DemoLogger) Fatal(args ...interface{}) {
	l.write(Fields{
		"lev": "F",
		"msg": fmt.Sprint(args...),
	})
	os.Exit(1)
}
func (l DemoLogger) Debugf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "D",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l DemoLogger) Infof(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "I",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l DemoLogger) Warnf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "W",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l DemoLogger) Errorf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "E",
		"msg": fmt.Sprintf(format, args...),
	})
}
func (l DemoLogger) Panicf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "P",
		"msg": fmt.Sprintf(format, args...),
	})
	panic(fmt.Sprintf(format, args...))
}
func (l DemoLogger) Fatalf(format string, args ...interface{}) {
	l.write(Fields{
		"lev": "F",
		"msg": fmt.Sprintf(format, args...),
	})
	os.Exit(1)
}
func (l DemoLogger) Debugw(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "D", msg
	l.write(fields)
}
func (l DemoLogger) Infow(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "I", msg
	l.write(fields)
}
func (l DemoLogger) Warnw(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "W", msg
	l.write(fields)
}
func (l DemoLogger) Errorw(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "E", msg
	l.write(fields)
}
func (l DemoLogger) Panicw(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "P", msg
	l.write(fields)
	panic(msg)
}
func (l DemoLogger) Fatalw(msg string, fields Fields) {
	fields["lev"], fields["msg"] = "F", msg
	l.write(fields)
	os.Exit(1)
}
func (l DemoLogger) With(fields Fields) Recorder {
	for key, val := range l.Fields {
		fields[key] = val
	}
	return DemoLogger{Out: l.Out, Name: l.Name, Fields: fields}
}
