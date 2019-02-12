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

func TestLogSrvAppender(t *testing.T) {
	var err error

	a := NewLogSrvAppender()
	a.Name = "h5_test"
	a.Addr = "localhost:50012"
	if a.Layout, err = ParseTemplateString("{time|20060102150405}|{level}|{message}"); err != nil {
		t.Fatal("ParseTemplate fail:", err)
	}
	if err = a.Open(); err != nil {
		t.Fatal("LogSrvAppender.Open fail:", err)
	}

	l := NewLogger()
	l.Level = LevelDebug
	l.AddAppender(a)

	l.Trace("tRaCe")
	l.Debug("dDbUg")
	l.Info("iNfO")
	l.Warn("WarN")
	l.Error("eRror")

	l.Close()
}
