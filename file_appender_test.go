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
	"os"
	"testing"
)

func TestFileAppender(t *testing.T) {
	const Path = "/tmp/log-go/file_appender_test.log"

	var err error

	os.Remove(Path) // remove existed file

	a := NewFileAppender()
	a.Path = Path
	if a.Layout, err = ParseTemplateString("{level}|{message}"); err != nil {
		t.Fatal("ParseTemplate fail:", err)
	}
	if err = a.Open(); err != nil {
		t.Fatal("FileAppender.Open fail:", err)
	}
	l := NewLogger()
	l.Level = LevelInfo
	l.AddAppender(a)

	l.Trace("tRaCe")
	l.Debug("dDbUg")
	l.Info("iNfO")
	l.Warn("WarN")
	l.Error("eRror")

	l.Close()

	b, err := ioutil.ReadFile(Path)
	if s1, s2 := string(b), "INFO|iNfO\nWARN|WarN\nERROR|eRror\n"; s1 != s2 {
		t.Fatalf("wrong file content: %q, expect: %q\n", s1, s2)
	}
}
