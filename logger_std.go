package log

import (
	"fmt"
	"log"
	"strings"
)

// StdLogger extend log package in standard library
type StdLogger struct {
	calldepth int

	kv []interface{}
}

func NewStdLogger(calldepth int) Logger {
	return &StdLogger{calldepth: calldepth}
}

func (l StdLogger) Output(
	lev Level, kvp []interface{}, v []interface{}) []interface{} {
	sb := strings.Builder{}
	{ // write level
		s := fmt.Sprintf("[%s] ", lev)
		if n := MaxLevelStringLen + 3 - len(s); n > 0 {
			s += strings.Repeat(" ", n)
		}
		sb.WriteString(s)
	}
	{ // write kv
		sb.WriteString("[")
		for i, s := range l.kv {
			if i > 0 {
				sb.WriteByte(' ')
			}
			sb.WriteString(s.(string))
		}
		for i, s := range l.fmtkvp(kvp) {
			if len(l.kv) > 0 || i > 0 {
				sb.WriteByte(' ')
			}
			sb.WriteString(s.(string))
		}
		sb.WriteString("] ")
	}
	{ // write v
		for _, v := range v {
			sb.WriteString(fmt.Sprintf("%v", v))
		}
	}
	log.Output(l.calldepth, sb.String())
	return v
}

func (l StdLogger) Debug(v ...interface{}) {
	l.Output(DebugLevel, nil, v)
}
func (l StdLogger) Info(v ...interface{}) {
	l.Output(InfoLevel, nil, v)
}
func (l StdLogger) Warn(v ...interface{}) {
	l.Output(WarnLevel, nil, v)
}
func (l StdLogger) Error(v ...interface{}) {
	l.Output(ErrorLevel, nil, v)
}
func (l StdLogger) Panic(v ...interface{}) {
	l.Output(PanicLevel, nil, v)
}
func (l StdLogger) Fatal(v ...interface{}) {
	l.Output(FatalLevel, nil, v)
}

func (l StdLogger) Debugf(format string, v ...interface{}) {
	l.Output(DebugLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}
func (l StdLogger) Infof(format string, v ...interface{}) {
	l.Output(InfoLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}
func (l StdLogger) Warnf(format string, v ...interface{}) {
	l.Output(WarnLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}
func (l StdLogger) Errorf(format string, v ...interface{}) {
	l.Output(ErrorLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}
func (l StdLogger) Panicf(format string, v ...interface{}) {
	l.Output(PanicLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}
func (l StdLogger) Fatalf(format string, v ...interface{}) {
	l.Output(FatalLevel, nil, []interface{}{fmt.Sprintf(format, v...)})
}

func (l StdLogger) Debugw(msg string, kvp ...interface{}) {
	l.Output(DebugLevel, kvp, []interface{}{msg})
}
func (l StdLogger) Infow(msg string, kvp ...interface{}) {
	l.Output(InfoLevel, kvp, []interface{}{msg})
}
func (l StdLogger) Warnw(msg string, kvp ...interface{}) {
	l.Output(WarnLevel, kvp, []interface{}{msg})
}
func (l StdLogger) Errorw(msg string, kvp ...interface{}) {
	l.Output(ErrorLevel, kvp, []interface{}{msg})
}
func (l StdLogger) Panicw(msg string, kvp ...interface{}) {
	l.Output(PanicLevel, kvp, []interface{}{msg})
}
func (l StdLogger) Fatalw(msg string, kvp ...interface{}) {
	l.Output(FatalLevel, kvp, []interface{}{msg})
}

func (l StdLogger) With(kvp ...interface{}) Logger {
	return StdLogger{
		calldepth: l.calldepth,
		kv:        append(l.kv, l.fmtkvp(kvp)...),
	}
}

func (StdLogger) fmtkvp(kvp []interface{}) []interface{} {
	v := make([]interface{}, 0, len(kvp)/2)
	for i := 0; i < len(kvp)-1; i += 2 {
		v = append(v, fmt.Sprintf("%v=%v", kvp[i], kvp[i+1]))
	}
	return v
}
