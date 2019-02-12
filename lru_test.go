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
	"testing"
)

func TestLru(t *testing.T) {
	l := newLru()
	l.touch("1")
	l.touch("2")
	l.touch("3")

	if s, _ := l.top(); s != "1" {
		t.Fatal(s)
	}
	l.pop()

	l.touch("2")
	if s, _ := l.top(); s != "3" {
		t.Fatal(s)
	}
	l.pop()

	l.touch("4")
	if s, _ := l.top(); s != "2" {
		t.Fatal(s)
	}
	l.pop()

	if s, _ := l.top(); s != "4" {
		t.Fatal(s)
	}
	l.pop()

	if !l.empty() {
		t.Fatal()
	}
}
