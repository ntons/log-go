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
	"time"
)

func TestTemplate(t *testing.T) {
	tm := time.Date(2019, 7, 22, 16, 33, 52, 23865, time.Local)

	e := Entry{
		FieldTime:    tm,
		FieldLevel:   LevelInfo,
		FieldMessage: "Hello World !@#$%^&*()",
	}

	input := []string{
		`/var/log/log-go/{time|20060102}/{time|15}.log`,
		`{time|2006-01-02 15:04:05}|{level}|{message}`,
	}
	output := []string{
		`/var/log/log-go/20190722/16.log`,
		`2019-07-22 16:33:52|INFO|Hello World !@#$%^&*()`,
	}

	for i, _ := range input {
		tpl, err := ParseTemplateString(input[i])
		if err != nil {
			t.Fatal("ParseTemplateString fail:", err)
		}
		if s := tpl.Format(e); s != output[i] {
			t.Fatalf("wrong format output: %q, expected: %q\n", s, output[i])
		}
	}
}

func TestTemplatePartialFormat(t *testing.T) {
	tm := time.Date(2019, 7, 22, 16, 33, 52, 23865, time.Local)

	e := Entry{
		FieldTime:    tm,
		FieldLevel:   LevelInfo,
		FieldMessage: "Hello World !@#$%^&*()",
	}

	input := `{time|060102} {level} {miss}`
	output := `190722 INFO {miss}`

	tpl, err := ParseTemplateString(input)
	if err != nil {
		t.Fatal("ParseTemplateString fail:", err)
	}
	if s := tpl.Format(e); s != output {
		t.Fatalf("wrong format output: %q, expected: %q\n", s, output)
	}
}
