package zap_test

import (
	"testing"

	_zap "go.uber.org/zap"

	"github.com/ntons/log-go/zap"
)

func TestZap(t *testing.T) {
	z, _ := _zap.NewDevelopment()

	l := zap.New(z, 1)

	l.Info("Hello", " World")

	l = l.With("foo", "bar")
	l.Infof("%s %s", "Hello", "World")

	l = l.With("bar", "bbbb")
	l.Infof("%s %s", "Hello", "World")
}
