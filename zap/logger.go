package zap

import (
	"github.com/ntons/log-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func fitFields(fields log.Fields) []interface{} {
	arr := make([]interface{}, 0, len(fields)*2)
	for key, val := range fields {
		arr = append(arr, key, val)
	}
	return arr
}

func fitLevel(lev log.Level) zapcore.Level {
	switch lev {
	case log.DebugLevel:
		return zapcore.DebugLevel
	case log.InfoLevel:
		return zapcore.InfoLevel
	case log.WarnLevel:
		return zapcore.WarnLevel
	case log.ErrorLevel:
		return zapcore.ErrorLevel
	case log.FatalLevel:
		return zapcore.FatalLevel
	default:
		return zapcore.PanicLevel
	}
}

type Recorder struct {
	sugar *zap.SugaredLogger
}

func (r *Recorder) With(fields log.Fields) log.Recorder {
	return &Recorder{sugar: r.sugar.With(fitFields(fields))}
}

func (r *Recorder) Debug(args ...interface{}) {
	r.sugar.Debug(args...)
}
func (r *Recorder) Info(args ...interface{}) {
	r.sugar.Info(args...)
}
func (r *Recorder) Warn(args ...interface{}) {
	r.sugar.Warn(args...)
}
func (r *Recorder) Error(args ...interface{}) {
	r.sugar.Error(args...)
}
func (r *Recorder) Fatal(args ...interface{}) {
	r.sugar.Fatal(args...)
}

func (r *Recorder) Debugf(format string, args ...interface{}) {
	r.sugar.Debugf(format, args...)
}
func (r *Recorder) Infof(format string, args ...interface{}) {
	r.sugar.Infof(format, args...)
}
func (r *Recorder) Warnf(format string, args ...interface{}) {
	r.sugar.Warnf(format, args...)
}
func (r *Recorder) Errorf(format string, args ...interface{}) {
	r.sugar.Errorf(format, args...)
}
func (r *Recorder) Fatalf(format string, args ...interface{}) {
	r.sugar.Fatalf(format, args...)
}

func (r *Recorder) Debugw(msg string, fields log.Fields) {
	r.sugar.Debugw(msg, fitFields(fields)...)
}
func (r *Recorder) Infow(msg string, fields log.Fields) {
	r.sugar.Infow(msg, fitFields(fields)...)
}
func (r *Recorder) Warnw(msg string, fields log.Fields) {
	r.sugar.Warnw(msg, fitFields(fields)...)
}
func (r *Recorder) Errorw(msg string, fields log.Fields) {
	r.sugar.Errorw(msg, fitFields(fields)...)
}
func (r *Recorder) Fatalw(msg string, fields log.Fields) {
	r.sugar.Fatalw(msg, fitFields(fields)...)
}

type Logger struct {
	*Recorder
	lev zap.AtomicLevel
}

func NewLogger(l *zap.Logger, lev zap.AtomicLevel) *Logger {
	return &Logger{
		Recorder: &Recorder{sugar: l.WithOptions(zap.AddCallerSkip(1)).Sugar()},
		lev:      lev,
	}
}

func (l *Logger) Close() error {
	return nil
}
func (l *Logger) Sync() error {
	return l.Recorder.sugar.Sync()
}
func (l *Logger) SetLevel(lev log.Level) {
	l.lev.SetLevel(fitLevel(lev))
}
func (l *Logger) IsLevelEnabled(lev log.Level) bool {
	return l.lev.Enabled(fitLevel(lev))
}

////////////////////////////////////////////////////////////////////////////////
// Builder registration
////////////////////////////////////////////////////////////////////////////////
type Builder struct {
	zap.Config
}

func NewBuilder() log.Builder {
	return &Builder{}
}
func NewProductionBuilder() log.Builder {
	return &Builder{zap.NewProductionConfig()}
}
func NewDevelopmentBuilder() log.Builder {
	return &Builder{zap.NewDevelopmentConfig()}
}
func (b *Builder) Build() (log.Logger, error) {
	if l, err := b.Config.Build(); err != nil {
		return nil, err
	} else {
		return NewLogger(l, b.Level), nil
	}
}
func init() {
	log.RegisterLogger("zap", NewBuilder)
	log.RegisterLogger("zap-pro", NewProductionBuilder)
	log.RegisterLogger("zap-dev", NewDevelopmentBuilder)
}
