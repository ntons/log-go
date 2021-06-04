package log

import (
	"fmt"
	"strings"
)

type Level int

const (
	_ Level = iota
	AllLevel
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

const MaxLevelStringLen = 5

func (lev Level) String() string {
	switch lev {
	case AllLevel:
		return "ALL"
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case PanicLevel:
		return "PANIC"
	case FatalLevel:
		return "FATAL"
	default:
		return ""
	}
}

func ParseLevel(s string) (Level, bool) {
	switch strings.ToUpper(s) {
	case "ALL":
		return AllLevel, true
	case "DEBUG":
		return DebugLevel, true
	case "INFO":
		return InfoLevel, true
	case "WARN":
		return WarnLevel, true
	case "ERROR":
		return ErrorLevel, true
	case "PANIC":
		return PanicLevel, true
	case "FATAL":
		return FatalLevel, true
	default:
		return 0, false
	}
}

func (lev *Level) UnmarshalText(b []byte) error {
	if len(b) > 0 {
		_lev, ok := ParseLevel(string(b))
		if !ok {
			return fmt.Errorf("invalid level: %s", b)
		}
		*lev = _lev
	}
	return nil
}
