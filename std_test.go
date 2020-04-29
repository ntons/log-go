package log

import (
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
		Debug(msg, Fields{"method": "Debugw"})
		Info(msg, Fields{"method": "Infow"})
		Warn(msg, Fields{"method": "Warnw"})
		Error(msg, Fields{"method": "Errorw"})
	}
	ReplaceStd(ConsoleLogger{}).Close()
	Log("TestStd")
	ReplaceStd(ConsoleLogger{}).Close()
	Log("TestStd")
	Close()
}
