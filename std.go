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
	"io/ioutil"
	"sync"
)

// global named loggers
type namedLogger struct {
	name   string
	logger *Logger
}

var (
	mu  sync.Mutex
	std *Logger = func() *Logger {
		s := "{time|2006/01/02 15:04:05} {level} {message}"
		appender := NewConsoleAppender()
		appender.Layout, _ = ParseTemplate([]byte(s))
		logger := NewLogger()
		logger.AddAppender(appender)
		return logger
	}()
)

func Close() {
	mu.Lock()
	defer mu.Unlock()
	if std != nil {
		std.Close()
		std = nil
	}
}

func ParseFile(path string) (err error) {
	var b []byte
	if b, err = ioutil.ReadFile(path); err != nil {
		return
	} else if err = ParseJSON(b); err != nil {
		return
	}
	return
}

func ParseJSON(b []byte) (err error) {
	logger := NewLogger()
	if err = logger.ParseJSON(b); err != nil {
		return
	}
	std = logger
	return
}

func Log(l Level, s string, a ...interface{}) {
	std.log(l, nil, s, a)
}
func Fatal(s string, a ...interface{}) {
	std.log(LevelFatal, nil, s, a)
}
func Error(s string, a ...interface{}) {
	std.log(LevelError, nil, s, a)
}
func Warn(s string, a ...interface{}) {
	std.log(LevelWarn, nil, s, a)
}
func Info(s string, a ...interface{}) {
	std.log(LevelInfo, nil, s, a)
}
func Debug(s string, a ...interface{}) {
	std.log(LevelDebug, nil, s, a)
}
func Trace(s string, a ...interface{}) {
	std.log(LevelTrace, nil, s, a)
}

func LogM(m M, l Level, s string, a ...interface{}) {
	std.log(l, m, s, a)
}
func FatalM(m M, s string, a ...interface{}) {
	std.log(LevelFatal, m, s, a)
}
func ErrorM(m M, s string, a ...interface{}) {
	std.log(LevelError, m, s, a)
}
func WarnM(m M, s string, a ...interface{}) {
	std.log(LevelWarn, m, s, a)
}
func InfoM(m M, s string, a ...interface{}) {
	std.log(LevelInfo, m, s, a)
}
func DebugM(m M, s string, a ...interface{}) {
	std.log(LevelDebug, m, s, a)
}
func TraceM(m M, s string, a ...interface{}) {
	std.log(LevelTrace, m, s, a)
}
