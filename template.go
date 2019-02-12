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
	"bytes"
	"errors"
	"fmt"
)

// formatter interface
type Formatter interface {
	Format(pat string) string
}

//
type TemplateVariable struct {
	Raw  string
	Args []string //[key, pat, dft]
}

//
type Template struct {
	Pat  string
	Vars []TemplateVariable
}

func ParseTemplate(b []byte) (t *Template, err error) {
	t = &Template{}
	if err = t.parse(append([]byte{}, b...)); err != nil {
		t = nil
		return
	}
	return
}
func ParseTemplateString(s string) (t *Template, err error) {
	t = &Template{}
	if err = t.parse([]byte(s)); err != nil {
		t = nil
		return
	}
	return
}

func (t *Template) parse(b []byte) (err error) {
	vars := make([]TemplateVariable, 0)
	for i, j := 0, len(b)-1; j >= 0; {
		if b[j] != '}' {
			j -= 1
			continue
		} else if j > 0 && b[j-1] == '}' {
			j -= 2
			continue
		}
		for i = j - 1; i >= 0; {
			if b[i] != '{' {
				i -= 1
				continue
			} else if i > 0 && b[i-1] == '{' {
				i -= 2
				continue
			}
			break
		}
		if b[i] != '{' {
			return errors.New(fmt.Sprintf(
				"invalid template expr at %d-%d, %s",
				i, j, b[i:j+1]))
		}
		v := TemplateVariable{Raw: string(b[i : j+1])}
		for _, a := range bytes.Split(b[i+1:j], []byte("|")) {
			v.Args = append(v.Args, string(a))
		}
		vars = append(vars, v)
		b = append(append(b[:i], '%', 's'), b[j+1:]...)
		j = i - 1
	}
	for i, l := 0, len(vars); i < l/2; i++ {
		vars[i], vars[l-1-i] = vars[l-1-i], vars[i]
	}
	t.Pat, t.Vars = string(b), vars
	return
}

func (t *Template) Format(e Entry) (s string) {
	a := make([]interface{}, 0, len(t.Vars))
	for _, _v := range t.Vars {
		var key, pat, dft string
		if l := len(_v.Args); l >= 3 {
			key, pat, dft = _v.Args[0], _v.Args[1], _v.Args[2]
		} else if l >= 2 {
			key, pat = _v.Args[0], _v.Args[1]
		} else if l >= 1 {
			key = _v.Args[0]
		}
		var s string
		if v, ok := e[key]; !ok {
			if dft != "" {
				s = dft
			} else {
				s = _v.Raw
			}
		} else if f, ok := v.(Formatter); ok {
			s = f.Format(pat)
		} else if pat != "" {
			s = fmt.Sprintf(pat, v)
		} else {
			s = fmt.Sprintf("%v", v)
		}
		a = append(a, s)
	}
	return fmt.Sprintf(t.Pat, a...)
}
