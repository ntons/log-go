package log

import (
	"strings"
)

type Level int

const (
	AllLevel Level = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
	UnknownLevel
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
		return AllLevel, false
	}
}
