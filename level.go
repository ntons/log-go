package log

import (
	"fmt"
	"strings"
)

type Level int8

// compare zap with logrus:
// logrus level panic is higher than fatal, however,
// I believe recoverable panic should be lower than fatal,
// just like what in zap.
// But, I don't judge it here, so it must noticed
// if level threshold was set to panic or fatal,
// whether it will be triggered depend on implement,
// although, the action after logging should be right.
const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel // panic after logging
	FatalLevel // os.Exit(1) after logging
)

func (l Level) String() (s string) {
	switch l {
	case DebugLevel:
		s = "debug"
	case InfoLevel:
		s = "info"
	case WarnLevel:
		s = "warn"
	case ErrorLevel:
		s = "error"
	case PanicLevel:
		s = "panic"
	case FatalLevel:
		s = "fatal"
	default:
		panic("invalid log level")
	}
	return
}

func (l Level) MarshalText() (p []byte, err error) {
	return []byte(l.String()), nil
}

func (l *Level) UnmarshalText(p []byte) (err error) {
	switch s := strings.ToLower(string(p)); s {
	case "debug":
		*l = DebugLevel
	case "info":
		*l = InfoLevel
	case "warn":
		*l = WarnLevel
	case "error":
		*l = ErrorLevel
	case "panic":
		*l = PanicLevel
	case "fatal":
		*l = FatalLevel
	default:
		err = fmt.Errorf("invalid log level %q", s)
	}
	return
}
