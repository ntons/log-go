package logrus

import (
	"runtime"
)

// Skip n Stacks from given frame
// 2 more stack should be skipped since this wrapper and log-go require
// Call on Entry.Caller if using a custom formatter.
// Call in callerPrettyfier if using builtin.
func CallerSkip(f *runtime.Frame, n int) *runtime.Frame {
	pcs := make([]uintptr, 25)
	depth := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for _f, more := frames.Next(); more; _f, more = frames.Next() {
		if _f.PC == f.PC {
			for i := 0; i < n && more; i++ {
				_f, more = frames.Next()
			}
			f = &_f
			break
		}
	}
	return f
}
