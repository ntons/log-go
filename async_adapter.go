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

// async appender wrapper, ONLY Write async

package log

import (
	"errors"
	"sync"
)

type AsyncAdapter struct {
	Appender
	//
	c chan Entry
	//
	wg   sync.WaitGroup
	quit chan struct{}
}

func NewAsyncAdapter(a Appender, bufferSize int) (x *AsyncAdapter) {
	if bufferSize <= 0 {
		bufferSize = 1000
	}
	x = &AsyncAdapter{
		Appender: a,
		c:        make(chan Entry, bufferSize),
		quit:     make(chan struct{}, 1),
	}
	// async write
	x.wg.Add(1)
	go func() {
		defer x.wg.Done()
		defer x.Appender.Close()

	loop:
		for {
			select {
			case <-x.quit:
				break loop
			case e := <-x.c:
				if err := x.Appender.Write(e); err != nil {
					x.Appender.OnFail(err)
				}
			}
		}
		for {
			select {
			case e := <-x.c:
				if err := x.Appender.Write(e); err != nil {
					x.Appender.OnFail(err)
				}
			default:
				return
			}
		}
	}()
	return
}

func (x *AsyncAdapter) Write(e Entry) (err error) {
	select {
	case x.c <- e:
	default:
		err = errors.New("async buffer full")
	}
	return
}

func (x *AsyncAdapter) Close() {
	if x.quit != nil {
		close(x.quit)
		x.wg.Wait()
		x.quit = nil
	}
}
