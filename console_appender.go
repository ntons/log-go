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
	"os"
)

func init() {
	GetAppenderFactory().Register(
		"ConsoleAppender",
		func() Appender { return NewConsoleAppender() })
}

type ConsoleAppender struct {
	*AppenderBase
}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{
		AppenderBase: NewAppenderBase(),
	}
}

func (a *ConsoleAppender) Open() (err error) {
	return
}

func (a *ConsoleAppender) Close() {
	os.Stdout.Sync()
}

func (a *ConsoleAppender) Write(e Entry) (err error) {
	if e.Level() < a.Level {
		return
	}
	_, err = os.Stdout.WriteString(a.format(e))
	return
}
