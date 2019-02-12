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
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestTemplatedRollingFileAppender(t *testing.T) {
	const Path = "/tmp/log-go/templated_rolling_file_appender_test_{level}_{time|150405}.log"
	const Prefix = "/tmp/log-go/templated_rolling_file_appender_test_"

	filepath.Walk(
		filepath.Dir(Path),
		func(s string, info os.FileInfo, err error) (_ error) {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return
			}
			if !strings.HasPrefix(s, Prefix) {
				return
			}
			os.Remove(s)
			return
		})

	var err error
	a := NewTemplatedRollingFileAppender()
	a.Path = Path
	if a.Layout, err = ParseTemplateString("{message}"); err != nil {
		t.Fatal("ParseTemplate fail:", err)
	}
	a.MaxFileSize = 100
	if err := a.Open(); err != nil {
		t.Fatal(err)
	}
	l := NewLogger()
	l.Level = LevelInfo
	l.AddAppender(a)

	for i := 0; i < 3; i++ {
		l.Info(strings.Repeat("1", 1000))
		l.Info(strings.Repeat("2", 24))
		l.Warn(strings.Repeat("3", 1024))
		l.Warn(strings.Repeat("3", 1024))
		l.Error(strings.Repeat("3", 1024))

		time.Sleep(time.Second)
	}

	l.Close()
}
