package logrus

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	"github.com/ntons/log-go"
	"github.com/sirupsen/logrus"
)

func fitFields(fields log.Fields) []interface{} {
	arr := make([]interface{}, 0, len(fields)*2)
	for key, val := range fields {
		arr = append(arr, key, val)
	}
	return arr
}

func fitLevel(lev log.Level) logrus.Level {
	switch lev {
	case log.DebugLevel:
		return logrus.DebugLevel
	case log.InfoLevel:
		return logrus.InfoLevel
	case log.WarnLevel:
		return logrus.WarnLevel
	case log.ErrorLevel:
		return logrus.ErrorLevel
	case log.FatalLevel:
		return logrus.FatalLevel
	default:
		return logrus.PanicLevel
	}
}

type Recorder struct {
	e *logrus.Entry
}

func (l *Recorder) With(fields log.Fields) log.Recorder {
	return &Recorder{
		e: l.e.WithFields(logrus.Fields(fields)),
	}
}

func (e *Recorder) Debug(args ...interface{}) {
	e.e.Debug(args...)
}
func (e *Recorder) Info(args ...interface{}) {
	e.e.Info(args...)
}
func (e *Recorder) Warn(args ...interface{}) {
	e.e.Warn(args...)
}
func (e *Recorder) Error(args ...interface{}) {
	e.e.Error(args...)
}
func (e *Recorder) Fatal(args ...interface{}) {
	e.e.Fatal(args...)
}

func (e *Recorder) Debugf(format string, args ...interface{}) {
	e.e.Debugf(format, args...)
}
func (e *Recorder) Infof(format string, args ...interface{}) {
	e.e.Infof(format, args...)
}
func (e *Recorder) Warnf(format string, args ...interface{}) {
	e.e.Warnf(format, args...)
}
func (e *Recorder) Errorf(format string, args ...interface{}) {
	e.e.Errorf(format, args...)
}
func (e *Recorder) Fatalf(format string, args ...interface{}) {
	e.e.Fatalf(format, args...)
}

func (e *Recorder) Debugw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Debug(msg)
}
func (e *Recorder) Infow(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Info(msg)
}
func (e *Recorder) Warnw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Warn(msg)
}
func (e *Recorder) Errorw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Error(msg)
}
func (e *Recorder) Fatalw(msg string, fields log.Fields) {
	e.e.WithFields(logrus.Fields(fields)).Fatal(msg)
}

type Logger struct {
	*Recorder

	outs []Out
}

func NewLogger(l *logrus.Logger) *Logger {
	e := logrus.NewEntry(l)
	return &Logger{Recorder: &Recorder{e: e}}
}
func (l *Logger) Sync() (err error) {
	for _, out := range l.outs {
		if syncErr := out.Sync(); syncErr != nil {
			err = syncErr
		}
	}
	return
}
func (l *Logger) Close() (err error) {
	for _, out := range l.outs {
		if closeErr := out.Close(); closeErr != nil {
			err = closeErr
		}
	}
	return
}

func (l *Logger) SetLevel(lev log.Level) {
	l.e.Logger.SetLevel(fitLevel(lev))
}
func (l *Logger) IsLevelEnabled(lev log.Level) bool {
	return l.e.Logger.IsLevelEnabled(fitLevel(lev))
}

///////////////////////////////////////////////////////////////////////////////
// Builder
///////////////////////////////////////////////////////////////////////////////
var (
	ErrBadFormatterConfig = errors.New("bad formatter config")
)

type FormatterConfig struct {
	Encoding string
}

type Builder struct {
	OutPaths     []string
	Formatter    json.RawMessage
	ReportCaller bool
	Level        logrus.Level
}

func newBuilder() log.Builder {
	return &Builder{}
}
func (cfg Builder) Build() (log.Logger, error) {
	outs, err := Open(cfg.OutPaths...)
	if err != nil {
		return nil, err
	}
	// build out
	writers := make([]io.Writer, 0, len(outs))
	for _, out := range outs {
		writers = append(writers, out)
	}
	// build formatter
	tok, err := json.NewDecoder(bytes.NewReader(cfg.Formatter)).Token()
	if err != nil {
		return nil, err
	}
	var builder FormatterBuilder
	switch tok.(type) {
	case string:
		if builder, err = newFormatterBuilder(tok.(string)); err != nil {
			return nil, err
		}
	case json.Delim:
		if tok.(json.Delim) != json.Delim('{') {
			return nil, ErrBadFormatterConfig
		}
		var fc FormatterConfig
		if err = json.Unmarshal(cfg.Formatter, &fc); err != nil {
			return nil, err
		}
		if builder, err = newFormatterBuilder(fc.Encoding); err != nil {
			return nil, err
		}
	default:
		return nil, ErrBadFormatterConfig
	}
	formatter, err := builder.Build()
	if err != nil {
		return nil, err
	}
	// build logger
	l := &logrus.Logger{
		Out:          io.MultiWriter(writers...),
		Formatter:    formatter,
		ReportCaller: cfg.ReportCaller,
		Level:        cfg.Level,
	}
	return &Logger{
		Recorder: &Recorder{e: logrus.NewEntry(l)},
		outs:     outs,
	}, nil
}

func init() {
	log.RegisterLogger("logrus", newBuilder)
}
