// Copyright 2019 The Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package log

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Logger struct {
	// 最低输出等级
	Level Level
	// 输出源
	appenders []Appender
	// 互斥量
	mu sync.Mutex
}

func NewLogger() *Logger {
	return &Logger{
		Level: LevelAll,
	}
}

func (x *Logger) ParseJSON(b []byte) (err error) {
	c := struct {
		Level     string
		Appenders []json.RawMessage
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if x.Level, err = ParseLevel(c.Level); err != nil {
		return
	}
	for _, b := range c.Appenders {
		var a Appender
		if a, err = GetAppenderFactory().Create(b); err != nil {
			return
		}
		x.AddAppender(a)
	}
	return
}

func (x *Logger) AddAppender(a Appender) {
	x.mu.Lock()
	defer x.mu.Unlock()
	for _, _a := range x.appenders {
		if _a == a {
			return
		}
	}
	x.appenders = append(x.appenders, a)
}
func (x *Logger) RemoveAppender(a Appender) {
	x.mu.Lock()
	defer x.mu.Unlock()
	for i, _a := range x.appenders {
		if _a == a {
			x.appenders = append(x.appenders[:i], x.appenders[i+1:]...)
			return
		}
	}
}

func (x *Logger) Close() {
	x.mu.Lock()
	defer x.mu.Unlock()
	for _, a := range x.appenders {
		a.Close()
	}
	x.appenders = x.appenders[:0]
}

type M map[string]interface{}

func (x *Logger) log(l Level, m M, s string, a []interface{}) {
	if l < x.Level {
		return
	}
	for i, v := range a {
		if f, ok := v.(func() string); ok {
			a[i] = f()
		}
	}

	// get file and line
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file, line = "unknown", 0
	} else {
		file = filepath.Base(file)
	}

	e := Entry{
		FieldTime:    time.Now(),
		FieldLevel:   l,
		FieldMessage: fmt.Sprintf(s, a...),
		FieldFile:    file,
		FieldLine:    line,
	}

	// add custom fields to entry
	for key, val := range m {
		if _, ok := e[key]; !ok {
			e[key] = val
		}
	}
	// write to all appenders
	x.mu.Lock()
	defer x.mu.Unlock()
	for _, a := range x.appenders {
		if err := a.Write(e); err != nil {
			a.OnFail(err)
			return
		}
	}
}

func (x *Logger) Log(l Level, s string, a ...interface{}) {
	x.log(l, nil, s, a)
}
func (x *Logger) Fatal(s string, a ...interface{}) {
	x.log(LevelFatal, nil, s, a)
}
func (x *Logger) Error(s string, a ...interface{}) {
	x.log(LevelError, nil, s, a)
}
func (x *Logger) Warn(s string, a ...interface{}) {
	x.log(LevelWarn, nil, s, a)
}
func (x *Logger) Info(s string, a ...interface{}) {
	x.log(LevelInfo, nil, s, a)
}
func (x *Logger) Debug(s string, a ...interface{}) {
	x.log(LevelDebug, nil, s, a)
}
func (x *Logger) Trace(s string, a ...interface{}) {
	x.log(LevelTrace, nil, s, a)
}

func (x *Logger) LogM(m M, l Level, s string, a ...interface{}) {
	x.log(l, m, s, a)
}
func (x *Logger) FatalM(m M, s string, a ...interface{}) {
	x.log(LevelFatal, m, s, a)
}
func (x *Logger) ErrorM(m M, s string, a ...interface{}) {
	x.log(LevelError, m, s, a)
}
func (x *Logger) WarnM(m M, s string, a ...interface{}) {
	x.log(LevelWarn, m, s, a)
}
func (x *Logger) InfoM(m M, s string, a ...interface{}) {
	x.log(LevelInfo, m, s, a)
}
func (x *Logger) DebugM(m M, s string, a ...interface{}) {
	x.log(LevelDebug, m, s, a)
}
func (x *Logger) TraceM(m M, s string, a ...interface{}) {
	x.log(LevelTrace, m, s, a)
}
