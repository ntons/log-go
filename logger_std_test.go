package log_test

import (
	_log "log"
	"testing"

	"github.com/ntons/log-go"
)

func TestStdLogger(t *testing.T) {
	_log.SetFlags(_log.Lshortfile)

	var l = log.NewStdLogger(3)

	l.Debug("Hello", " World")
	l.Info("Hello", " World")
	l.Warn("Hello", " World")
	l.Error("Hello", " World")

	l.Debugf("%s %s", "Hello", "World")
	l.Infof("%s %s", "Hello", "World")
	l.Warnf("%s %s", "Hello", "World")
	l.Errorf("%s %s", "Hello", "World")

	l.Debugw("Hello World", "foo", "bar")
	l.Infow("Hello World", "foo", "bar")
	l.Warnw("Hello World", "foo", "bar")
	l.Errorw("Hello World", "foo", "bar")

	l = l.With("foo", "bar")

	l.Debug("Hello", " World")
	l.Info("Hello", " World")
	l.Warn("Hello", " World")
	l.Error("Hello", " World")

	l.Debugf("%s %s", "Hello", "World")
	l.Infof("%s %s", "Hello", "World")
	l.Warnf("%s %s", "Hello", "World")
	l.Errorf("%s %s", "Hello", "World")

	l.Debugw("Hello World", "foo1", "bar1")
	l.Infow("Hello World", "foo1", "bar1")
	l.Warnw("Hello World", "foo1", "bar1")
	l.Errorw("Hello World", "foo1", "bar1")
}
