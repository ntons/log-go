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
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// level
type Level int

const (
	LevelAll   Level = 0
	LevelTrace Level = 100
	LevelDebug Level = 200
	LevelInfo  Level = 300
	LevelWarn  Level = 400
	LevelError Level = 500
	LevelFatal Level = 600
)

var (
	levStrMap = map[Level]string{
		LevelAll:   "ALL",
		LevelTrace: "TRACE",
		LevelDebug: "DEBUG",
		LevelInfo:  "INFO",
		LevelWarn:  "WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
	}
	strLevMap = map[string]Level{
		"ALL":   LevelAll,
		"TRACE": LevelTrace,
		"DEBUG": LevelDebug,
		"INFO":  LevelInfo,
		"WARN":  LevelWarn,
		"ERROR": LevelError,
		"FATAL": LevelFatal,
	}
)

func RegisterLevel(l Level, s string) (err error) {
	if _, ok := levStrMap[l]; ok {
		return errors.New("duplicate level value")
	}
	if _, ok := strLevMap[s]; ok {
		return errors.New("duplicate level name")
	}
	levStrMap[l] = s
	strLevMap[s] = l
	return
}

func ParseLevel(s string) (l Level, err error) {
	var ok bool
	if l, ok = strLevMap[strings.ToUpper(s)]; !ok {
		var v int
		if v, err = strconv.Atoi(s); err != nil {
			err = errors.New("invalid syntax: " + s)
		} else {
			l = Level(v)
		}
	}
	return
}

func (l Level) String() string {
	if s, ok := levStrMap[l]; ok {
		return s
	} else {
		return fmt.Sprintf("%d", l)
	}
}

// log entry
const (
	FieldTime    = "time"
	FieldLevel   = "level"
	FieldFile    = "file"
	FieldLine    = "line"
	FieldMessage = "message"
)

type Entry map[string]interface{}

// panic if not set
func (e Entry) Level() Level {
	return e[FieldLevel].(Level)
}
func (e Entry) Time() time.Time {
	return e[FieldTime].(time.Time)
}
