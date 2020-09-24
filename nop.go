package log

type Nop struct{}

func (Nop) Debug(args ...interface{}) {}
func (Nop) Info(args ...interface{})  {}
func (Nop) Warn(args ...interface{})  {}
func (Nop) Error(args ...interface{}) {}
func (Nop) Panic(args ...interface{}) {}
func (Nop) Fatal(args ...interface{}) {}

func (Nop) Debugf(format string, args ...interface{}) {}
func (Nop) Infof(format string, args ...interface{})  {}
func (Nop) Warnf(format string, args ...interface{})  {}
func (Nop) Errorf(format string, args ...interface{}) {}
func (Nop) Panicf(format string, args ...interface{}) {}
func (Nop) Fatalf(format string, args ...interface{}) {}

func (Nop) Debugw(msg string, keyValuePairs ...interface{}) {}
func (Nop) Infow(msg string, keyValuePairs ...interface{})  {}
func (Nop) Warnw(msg string, keyValuePairs ...interface{})  {}
func (Nop) Errorw(msg string, keyValuePairs ...interface{}) {}
func (Nop) Panicw(msg string, keyValuePairs ...interface{}) {}
func (Nop) Fatalw(msg string, keyValuePairs ...interface{}) {}

func (Nop) With(fields Fields) Recorder                 { return Nop{} }
func (Nop) Withw(keyValuePairs ...interface{}) Recorder { return Nop{} }
