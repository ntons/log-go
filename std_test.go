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
		Debugw(msg, "method", "Debugw")
		Infow(msg, "method", "Infow")
		Warnw(msg, "method", "Warnw")
		Errorw(msg, "method", "Errorw")
	}
	SwapStd(NewConsoleLogger()).Close()
	Log("TestStd")
	SwapStd(NewConsoleLogger()).Close()
	Log("TestStd")
	Close()
}
