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
	"os"
	"time"
)

func init() {
	GetAppenderFactory().Register(
		"TemplatedRollingFileAppender",
		func() Appender { return NewTemplatedRollingFileAppender() })
}

// TemplatedRollingFileAppender provides both size-based rolling
// and pattern-based rolling
type TemplatedRollingFileAppender struct {
	*RollingFileAppender

	// path template
	tPath *Template

	// max open file, if exceeded, lru works
	MaxOpenFile int
	// max life time of inactive file, if exceeded, close
	MaxIdleTime time.Duration

	// files
	fileMap map[string]*os.File
	fileLru *lru
}

func NewTemplatedRollingFileAppender() *TemplatedRollingFileAppender {
	return &TemplatedRollingFileAppender{
		RollingFileAppender: NewRollingFileAppender(),
		MaxOpenFile:         100,
		MaxIdleTime:         time.Hour,
		fileMap:             make(map[string]*os.File),
		fileLru:             newLru(),
	}
}

func (a *TemplatedRollingFileAppender) ParseJSON(b []byte) (err error) {
	if err = a.RollingFileAppender.ParseJSON(b); err != nil {
		return
	}
	c := struct {
		MaxOpenFile int
		MaxIdleTime string
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	a.MaxOpenFile = c.MaxOpenFile
	if c.MaxIdleTime != "" {
		var d time.Duration
		if d, err = time.ParseDuration(c.MaxIdleTime); err != nil {
			return
		} else {
			a.MaxIdleTime = d
		}
	}
	return
}

func (a *TemplatedRollingFileAppender) Open() (err error) {
	if a.Path == "" {
		return errors.New("file path is missing")
	}
	if a.tPath, err = ParseTemplateString(a.Path); err != nil {
		return
	}
	return // open on Write
}

func (a *TemplatedRollingFileAppender) Close() {
	for _, file := range a.fileMap {
		file.Sync()
		file.Close()
	}
	a.fileMap = make(map[string]*os.File)
}

func (a *TemplatedRollingFileAppender) Write(e Entry) (err error) {
	if e.Level() < a.Level {
		return
	}
	if err = a.rollover(e); err != nil {
		return
	}
	if err = a.RollingFileAppender.Write(e); err != nil {
		return
	}
	a.fileMap[a.Path] = a.file // file may be changed on write
	a.fileLru.touch(a.Path)
	a.lru()
	return
}

func (a *TemplatedRollingFileAppender) rollover(e Entry) (err error) {
	if path := a.tPath.Format(e); a.file == nil || path != a.Path {
		a.Path = path
		if file, ok := a.fileMap[a.Path]; ok {
			a.file = file
		} else {
			if err = a.RollingFileAppender.Open(); err != nil {
				return
			}
			a.fileMap[a.Path] = a.file
			a.fileLru.touch(a.Path)
		}
	}
	return
}

func (a *TemplatedRollingFileAppender) lru() {
	for !a.fileLru.empty() {
		s, t := a.fileLru.top()
		if !((a.MaxIdleTime > 0 && time.Since(t) > a.MaxIdleTime) ||
			(a.MaxOpenFile > 0 && len(a.fileMap) > a.MaxOpenFile)) {
			break
		}
		a.fileLru.pop()
		file := a.fileMap[s]
		delete(a.fileMap, s)
		file.Sync()
		file.Close()
	}
}
