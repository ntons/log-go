package logrus

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

type FormatterBuilder interface {
	Build() (logrus.Formatter, error)
}

var (
	formatterFactoryMtx sync.Mutex
	formatterFactoryMap = make(map[string]func() FormatterBuilder)
)

func RegisterFormatter(name string, factory func() FormatterBuilder) error {
	name = strings.ToLower(name)
	formatterFactoryMtx.Lock()
	defer formatterFactoryMtx.Unlock()
	if _, ok := formatterFactoryMap[name]; ok {
		return fmt.Errorf("formatter already registered for name %q", name)
	}
	formatterFactoryMap[name] = factory
	return nil
}

func newFormatterBuilder(name string) (FormatterBuilder, error) {
	name = strings.ToLower(name)
	formatterFactoryMtx.Lock()
	defer formatterFactoryMtx.Unlock()
	factory, ok := formatterFactoryMap[name]
	if !ok {
		return nil, fmt.Errorf("formatter not been registered for name %q", name)
	}
	return factory(), nil
}

func callerPrettyfier(f *runtime.Frame) (function string, file string) {
	// skip one more stack，f must in stack here
	pcs := make([]uintptr, 25)
	depth := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for _f, more := frames.Next(); more; _f, more = frames.Next() {
		if _f.PC == f.PC {
			for i := 0; i < 2 && more; i++ { // skip 2 more stack
				_f, more = frames.Next()
			}
			f = &_f
			break
		}
	}
	// prettify
	file = fmt.Sprintf("%s:%d", f.File, f.Line)
	if i := strings.LastIndexByte(file, '/'); i > 0 {
		if i = strings.LastIndexByte(file[:i], '/'); i > 0 {
			file = file[i+1:]
		}
	}
	return f.Function, file
}

type textFormatterBuilder struct {
	*logrus.TextFormatter
}

func newTextFormatterBuilder() FormatterBuilder {
	return &textFormatterBuilder{
		TextFormatter: &logrus.TextFormatter{
			CallerPrettyfier: callerPrettyfier,
		},
	}
}
func (f textFormatterBuilder) Build() (logrus.Formatter, error) {
	return f.TextFormatter, nil
}

type jsonFormatterBuilder struct {
	*logrus.JSONFormatter
}

func newJSONFormatterBuilder() FormatterBuilder {
	return &jsonFormatterBuilder{
		JSONFormatter: &logrus.JSONFormatter{
			CallerPrettyfier: callerPrettyfier,
		},
	}
}
func (f jsonFormatterBuilder) Build() (logrus.Formatter, error) {
	return f.JSONFormatter, nil
}

func init() {
	RegisterFormatter("text", newTextFormatterBuilder)
	RegisterFormatter("json", newJSONFormatterBuilder)
}
