package log

import (
	"testing"
)

func TestIsNop(t *testing.T) {
	if !IsNop(nop) {
		t.Fatal("fail")
	}
	if !IsNop(nopLogger{}) {
		t.Fatal("fail")
	}
	if !IsNop(&nopLogger{}) {
		t.Fatal("fail")
	}
}
