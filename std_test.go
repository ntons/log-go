package log

import (
	"testing"
)

func TestStd(t *testing.T) {
	Debugf("xxx")
	Infof("xxx")
	r := Withw("foo", "bar")
	r.Infof("xxx")
	r = r.Withw("aaa", "bbb")
	r.Infof("xxx")
}
