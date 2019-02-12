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
	"container/list"
	"time"
)

type lruValue struct {
	s string
	t time.Time
}

// 同时兼顾heap特征和查找效率，只能用排序数组
type lru struct {
	// value list
	vlist *list.List
	// s->elem index
	index map[string]*list.Element
}

func newLru() *lru {
	return &lru{
		vlist: list.New(),
		index: make(map[string]*list.Element),
	}
}

func (l *lru) len() int {
	return l.vlist.Len()
}
func (l *lru) empty() bool {
	return l.vlist.Len() == 0
}

func (l *lru) pop() {
	if e := l.vlist.Front(); e != nil {
		delete(l.index, e.Value.(*lruValue).s)
		l.vlist.Remove(e)
	}
}
func (l *lru) top() (s string, t time.Time) {
	if e := l.vlist.Front(); e != nil {
		v := e.Value.(*lruValue)
		s, t = v.s, v.t
	}
	return
}
func (l *lru) touch(s string) {
	if s == "" {
		return
	}
	if e, ok := l.index[s]; ok {
		e.Value.(*lruValue).t = time.Now()
		l.vlist.MoveToBack(e)
	} else {
		l.vlist.PushBack(&lruValue{s: s, t: time.Now()})
		l.index[s] = l.vlist.Back()
	}
}
