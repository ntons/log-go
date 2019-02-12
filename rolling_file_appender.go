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
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func init() {
	GetAppenderFactory().Register(
		"RollingFileAppender",
		func() Appender { return NewRollingFileAppender() })
}

// rolling by size
type RollingFileAppender struct {
	*FileAppender
	//
	MaxFileSize int
}

func NewRollingFileAppender() *RollingFileAppender {
	return &RollingFileAppender{
		FileAppender: NewFileAppender(),
	}
}

func (a *RollingFileAppender) ParseJSON(b []byte) (err error) {
	if a.FileAppender.ParseJSON(b); err != nil {
		return
	}
	c := struct{ MaxFileSize string }{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if s := strings.TrimRight(c.MaxFileSize, "B"); s != "" {
		n := 1
		if strings.HasSuffix(s, "k") || strings.HasSuffix(s, "K") {
			s, n = s[:len(s)-1], 1024
		} else if strings.HasSuffix(s, "m") || strings.HasSuffix(s, "M") {
			s, n = s[:len(s)-1], 1024*1024
		} else if strings.HasSuffix(s, "g") || strings.HasSuffix(s, "G") {
			s, n = s[:len(s)-1], 1024*1024*1024
		}
		v := 0
		if v, err = strconv.Atoi(s); err != nil {
			err = errors.New("bad max file size")
			return
		}
		a.MaxFileSize = v * n
	}
	return
}

func (a *RollingFileAppender) Write(e Entry) (err error) {
	if e.Level() < a.Level {
		return
	}
	if err = a.rollover(); err != nil {
		return
	}
	if err = a.FileAppender.Write(e); err != nil {
		return
	}
	return
}

func (a *RollingFileAppender) rollover() (err error) {
	if a.file != nil && a.MaxFileSize > 0 {
		var info os.FileInfo
		if info, err = a.file.Stat(); err != nil {
			return
		}
		if int(info.Size()) >= a.MaxFileSize {
			a.FileAppender.Close()
			backFilePath := ""
			if backFilePath, err = a.backFilePath(); err != nil {
				return
			}
			os.Rename(a.Path, backFilePath)
			if a.FileAppender.Open(); err != nil {
				return
			}
		}
	}
	return
}

func (a *RollingFileAppender) backFilePath() (path string, err error) {
	ilist := make([]int, 0)
	if err = filepath.Walk(
		filepath.Dir(a.Path),
		func(path string, info os.FileInfo, err error) (_ error) {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return
			}
			s := a.Path + "."
			if !strings.HasPrefix(path, s) {
				return
			}
			i, err := strconv.Atoi(path[len(s):])
			if err != nil {
				return
			}
			ilist = append(ilist, i)
			return
		}); err != nil {
		return
	}
	i := 0
	if n := len(ilist); n > 0 {
		sort.Ints(ilist)
		i = ilist[n-1]
	}
	path = fmt.Sprintf("%s.%d", a.Path, i+1)
	return
}
