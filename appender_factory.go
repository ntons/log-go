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
	"sync"
)

type AppenderCreateFunc func() Appender

type AppenderFactory struct {
	m map[string]AppenderCreateFunc
}

func (x *AppenderFactory) Register(
	name string, creator AppenderCreateFunc) (err error) {
	if _, exist := x.m[name]; exist {
		return errors.New("appender had been registered")
	}
	x.m[name] = creator
	return
}

func (x *AppenderFactory) Create(b []byte) (a Appender, err error) {
	c := struct {
		Type            string
		Async           bool
		AsyncBufferSize int
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	creator, ok := x.m[c.Type]
	if !ok {
		err = errors.New(fmt.Sprintf("unknown appender: %s", c.Type))
		return
	}
	a = creator()
	if err = a.ParseJSON(b); err != nil {
		return
	}
	if c.Async {
		a = NewAsyncAdapter(a, c.AsyncBufferSize)
	}
	if err = a.Open(); err != nil {
	}
	return
}

var (
	// don't use these directly
	__AppenderFactoryInstance     *AppenderFactory
	__AppenderFactoryInstanceOnce sync.Once
)

// Singleton
func GetAppenderFactory() *AppenderFactory {
	__AppenderFactoryInstanceOnce.Do(func() {
		__AppenderFactoryInstance = &AppenderFactory{
			m: make(map[string]AppenderCreateFunc),
		}
	})
	return __AppenderFactoryInstance
}
