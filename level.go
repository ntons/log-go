package log

import (
	"fmt"
	"strings"
)

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
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
	case FatalLevel:
		s = "fatal"
	default:
		s = fmt.Sprintf("%d", l)
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
	case "fatal":
		*l = FatalLevel
	default:
		err = fmt.Errorf("invalid level %q", s)
	}
	return
}
