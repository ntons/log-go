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
	"errors"
	"fmt"
	"os"
)

type Appender interface {
	// parse configure
	ParseJSON(b []byte) (err error)
	// open to write
	Open() (err error)
	// close appender
	Close()
	// write log entry
	Write(e Entry) (err error)
	// on fail callback
	OnFail(err error)
}

type AppenderBase struct {
	Level  Level
	Layout *Template

	failFunc func(err error)
}

func NewAppenderBase() *AppenderBase {
	return &AppenderBase{}
}

func (a *AppenderBase) ParseJSON(b []byte) (err error) {
	c := struct {
		Level  string
		Layout []byte
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if c.Level == "" {
		return errors.New("missing level")
	}
	if len(c.Layout) == 0 {
		return errors.New("missing layout")
	}
	if a.Level, err = ParseLevel(c.Level); err != nil {
		return
	}
	if a.Layout, err = ParseTemplate(c.Layout); err != nil {
		return
	}
	return
}

func (a *AppenderBase) OnFail(err error) {
	fmt.Fprintf(os.Stderr, "failed to log: %s\n", err)
	if a.failFunc != nil {
		a.failFunc(err)
	}
}

func (a *AppenderBase) format(e Entry) (s string) {
	return a.Layout.Format(e) + "\n"
}
