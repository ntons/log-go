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
	"strconv"
	"syscall"
)

func init() {
	GetAppenderFactory().Register(
		"FileAppender",
		func() Appender { return NewFileAppender() })
}

type FileAppender struct {
	*AppenderBase
	// required options
	Path string
	// optional options
	DirPerm   os.FileMode
	FilePerm  os.FileMode
	SyncWrite bool
	// 文件对象
	file *os.File
}

func NewFileAppender() *FileAppender {
	return &FileAppender{
		AppenderBase: NewAppenderBase(),
		DirPerm:      0755,
		FilePerm:     0644,
		SyncWrite:    false,
	}
}

func (a *FileAppender) ParseJSON(b []byte) (err error) {
	if err = a.AppenderBase.ParseJSON(b); err != nil {
		return
	}
	c := struct {
		Path      string
		DirPerm   string
		FilePerm  string
		SyncWrite bool
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if c.Path == "" {
		err = errors.New("file path required")
		return
	} else {
		a.Path = c.Path
	}
	if c.DirPerm != "" {
		var v uint64
		if v, err = strconv.ParseUint(c.DirPerm, 8, 16); err != nil {
			err = errors.New("bad dir perm")
			return
		}
		a.DirPerm = os.FileMode(v)
	}
	if c.FilePerm != "" {
		var v uint64
		if v, err = strconv.ParseUint(c.FilePerm, 8, 16); err != nil {
			err = errors.New("bad file perm")
			return
		}
		a.FilePerm = os.FileMode(v)
	}
	a.SyncWrite = c.SyncWrite
	return
}

func (a *FileAppender) Open() (err error) {
	if a.Path == "" {
		return errors.New("file path is missing")
	}
	if err = os.MkdirAll(filepath.Dir(a.Path), a.DirPerm); err != nil {
		return
	}
	var (
		file *os.File
		flag = os.O_WRONLY | os.O_APPEND | os.O_CREATE
	)
	if a.SyncWrite {
		flag |= os.O_SYNC
	}
	if file, err = os.OpenFile(a.Path, flag, a.FilePerm); err != nil {
		return errors.New(fmt.Sprintf("open file fail: %v", err))
	}
	if err = syscall.Flock(
		int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		file.Close()
		return errors.New(fmt.Sprintf("flock fail: %v", err))
	}
	a.file = file
	return
}

func (a *FileAppender) Close() {
	if a.file != nil {
		a.file.Sync()
		a.file.Close()
		a.file = nil
	}
}

func (a *FileAppender) Write(e Entry) (err error) {
	if e.Level() < a.Level {
		return
	}
	if a.file == nil {
		err = errors.New("file not open")
		return
	}
	if _, err = a.file.WriteString(a.format(e)); err != nil {
		return
	}
	return
}
