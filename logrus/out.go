package logrus

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"sync"
)

type Syncer interface {
	Sync() error
}

type Out interface {
	io.WriteCloser
	Syncer
}

type nopCloseOut struct {
	Out
}

func nopClose(out Out) Out {
	return nopCloseOut{out}
}
func (nopCloseOut) Close() error {
	return nil
}

var (
	outFactoryMtx sync.Mutex
	outFactoryMap = make(map[string]func(*url.URL) (Out, error))
)

func RegisterOut(name string, factory func(*url.URL) (Out, error)) error {
	outFactoryMtx.Lock()
	defer outFactoryMtx.Unlock()
	if _, ok := outFactoryMap[name]; ok {
		return fmt.Errorf("out already registered for name %q", name)
	}
	outFactoryMap[name] = factory
	return nil
}

func newOut(path string) (Out, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("can't parse %q as a URL: %v", path, err)
	}
	outFactoryMtx.Lock()
	defer outFactoryMtx.Unlock()
	t := u.Scheme
	if t == "" {
		t = u.Path // for stderr and stdout
	}
	factory, ok := outFactoryMap[t]
	if !ok {
		return nil, fmt.Errorf("no out found for scheme %q", t)
	}
	return factory(u)
}

type stdWrapper struct {
	Out
}

func wrapStd(out Out) Out {
	return stdWrapper{Out: out}
}
func (w stdWrapper) Close() error {
	return w.Sync()
}

var (
	wrappedStderr = wrapStd(os.Stderr)
	wrappedStdout = wrapStd(os.Stdout)
)

func init() {
	RegisterOut("stdout", func(*url.URL) (Out, error) {
		return wrappedStdout, nil
	})
	RegisterOut("stderr", func(*url.URL) (Out, error) {
		return wrappedStderr, nil
	})
}

func Open(paths ...string) (_ []Out, err error) {
	outs := make([]Out, 0, len(paths))
	for _, path := range paths {
		var out Out
		if out, err = newOut(path); err != nil {
			break
		}
		outs = append(outs, out)
	}
	if err != nil {
		for _, out := range outs {
			out.Close()
		}
		return
	}
	return outs, nil
}
