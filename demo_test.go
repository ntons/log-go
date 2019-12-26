package log

import (
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	var l Logger = NewDemoLogger(os.Stdout, "demo")
	Log := func(r Recorder, msg string) {
		r.Debug(msg, " Debug")
		r.Info(msg, " Info")
		r.Warn(msg, " Warn")
		r.Error(msg, " Error")
		r.Debugf("%v %v", msg, "Debugf")
		r.Infof("%v %v", msg, "Infof")
		r.Warnf("%v %v", msg, "Warnf")
		r.Errorf("%v %v", msg, "Errorf")
		r.Debugw(msg, Fields{"method": "Debugw"})
		r.Infow(msg, Fields{"method": "Infow"})
		r.Warnw(msg, Fields{"method": "Warnw"})
		r.Errorw(msg, Fields{"method": "Errorw"})
	}
	Log(l, "logger")
	Log(l.With(Fields{"foo": "bar"}), "with")
}
