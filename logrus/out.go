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
	factory, ok := outFactoryMap[u.Scheme]
	if !ok {
		return nil, fmt.Errorf("no sink found for scheme %q", u.Scheme)
	}
	return factory(u)
}

func init() {
	RegisterOut("stdout", func(*url.URL) (Out, error) {
		return os.Stdout, nil
	})
	RegisterOut("stderr", func(*url.URL) (Out, error) {
		return os.Stderr, nil
	})
}

func Open(paths ...string) ([]Out, error) {
	success := false
	outs := make([]Out, 0, len(paths))
	defer func() {
		if !success {
			for _, out := range outs {
				out.Close()
			}
		}
	}()
	for _, path := range paths {
		out, err := newOut(path)
		if err != nil {
			return nil, err
		}
		outs = append(outs, out)
	}
	success = true
	return outs, nil
}
