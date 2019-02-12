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
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/ntons/log-go/proto"
	"google.golang.org/grpc"
)

func init() {
	GetAppenderFactory().Register(
		"LogSrvAppender",
		func() Appender { return NewLogSrvAppender() },
	)
}

type LogSrvAppender struct {
	*AppenderBase
	Name string
	Addr string

	conn *grpc.ClientConn
	cli  proto.LogClient
}

func NewLogSrvAppender() *LogSrvAppender {
	return &LogSrvAppender{
		AppenderBase: NewAppenderBase(),
	}
}

func (a *LogSrvAppender) ParseJSON(b []byte) (err error) {
	c := struct {
		Name string
		Addr string
	}{}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if c.Name == "" || c.Addr == "" {
		return errors.New("option(s) missed")
	}
	return
}

func (a *LogSrvAppender) Open() (err error) {
	if a.conn, err = grpc.Dial(a.Addr, grpc.WithInsecure()); err != nil {
		return
	}
	a.cli = proto.NewLogClient(a.conn)
	return
}

func (a *LogSrvAppender) Close() {
	if a.conn != nil {
		a.conn.Close()
		a.conn = nil
		a.cli = nil
	}
}

func (a *LogSrvAppender) Write(e Entry) (err error) {
	if e.Level() < a.Level {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err = a.cli.Write(ctx, &proto.LogReq{
		Name:  a.Name,
		Time:  e.Time().UnixNano(),
		Level: int64(e.Level()),
		Text:  a.format(e),
	}); err != nil {
		return
	}
	return
}
