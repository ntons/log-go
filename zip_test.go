package log

import (
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	ReplaceStd(NewDemoLogger(os.Stdout, "log1")).Close()
	w1 := With(Fields{"foo1": "bar"})
	w2 := w1.With(Fields{"foo2": "bar"})
	w3 := w2.With(Fields{"foo3": "bar"})
	w4 := w2.With(Fields{"foo4": "bar"})
	w1.Info("w1 before replace")
	w2.Info("w2 before replace")
	w3.Info("w3 before replace")
	w4.Info("w4 before replace")
	ReplaceStd(NewDemoLogger(os.Stdout, "log2")).Close()
	w3.Info("w3 after replace")
	w4.Info("w4 after replace")
	w1.Info("w1 after replace")
	w2.Info("w2 after replace")
}
