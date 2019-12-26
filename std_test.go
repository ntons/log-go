package log

import (
	"os"
	"testing"
)

func TestStd(t *testing.T) {
	Log := func(msg string) {
		Debug(msg, " Debug")
		Info(msg, " Info")
		Warn(msg, " Warn")
		Error(msg, " Error")
		Debugf("%v %v", msg, "Debugf")
		Infof("%v %v", msg, "Infof")
		Warnf("%v %v", msg, "Warnf")
		Errorf("%v %v", msg, "Errorf")
		Debugw(msg, Fields{"method": "Debugw"})
		Infow(msg, Fields{"method": "Infow"})
		Warnw(msg, Fields{"method": "Warnw"})
		Errorw(msg, Fields{"method": "Errorw"})
	}
	ReplaceStd(NewDemoLogger(os.Stdout, "log1")).Close()
	Log("TestStd")
	ReplaceStd(NewDemoLogger(os.Stdout, "log2")).Close()
	Log("TestStd")
	Close()
}
