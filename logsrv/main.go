package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/ntons/log-go"
	"github.com/ntons/log-go/proto"
)

var (
	OptAddr   string
	OptPath   string
	OptLevel  = log.LevelAll
	OptConfig string
)

type Srv struct {
	a *log.TemplatedRollingFileAppender
}

func NewSrv() *Srv {
	return &Srv{}
}

func (s *Srv) Init() (err error) {
	a := log.NewTemplatedRollingFileAppender()
	a.Layout, _ = log.ParseTemplateString("{text|%s}")
	a.Path = OptPath
	a.Level = OptLevel
	if err = a.Open(); err != nil {
		return
	}
	s.a = a
	return
}

func (s *Srv) Close() {
	if s.a != nil {
		s.a.Close()
		s.a = nil
	}
}

func (s *Srv) Write(
	ctx context.Context, req *proto.LogReq) (rep *proto.LogRep, err error) {
	e := log.Entry{
		log.FieldTime:  time.Unix(0, req.Time),
		log.FieldLevel: log.Level(req.Level),
		"name":         req.Name,
		"text":         req.Text,
	}
	if err = s.a.Write(e); err != nil {
		return
	}
	rep = &proto.LogRep{}
	return
}

func ParseConfig() (err error) {
	c := struct {
		Addr  string
		Path  string
		Level string
	}{}
	b, err := ioutil.ReadFile(OptConfig)
	if err != nil {
		return
	}
	if err = json.Unmarshal(b, &c); err != nil {
		return
	}
	if OptAddr == "" {
		OptAddr = c.Addr
	}
	if OptPath == "" {
		OptPath = c.Path
	}
	if OptLevel == log.LevelAll && c.Level != "" {
		if OptLevel, err = log.ParseLevel(c.Level); err != nil {
			err = errors.New("syntex error: -level")
			return
		}
	}
	return
}

func PrintUsageToStderr() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

func Init() (err error) {
	var (
		OptLevelStr string
	)
	// parse options
	flag.StringVar(&OptAddr, "a", "",
		"[REQUIRED] Listen address")
	flag.StringVar(&OptPath, "p", "",
		"[REQUIRED] Path template(name,time,level)")
	flag.StringVar(&OptLevelStr, "l", "all",
		"[OPTIONAL] Log level")
	flag.StringVar(&OptConfig, "c", "",
		"[OPTIONAL] Config file")
	flag.Parse()

	if OptConfig != "" {
		if err = ParseConfig(); err != nil {
			return
		}
	}
	if OptAddr == "" {
		return errors.New("flag is required: -addr")
	}
	if OptPath == "" {
		return errors.New("flag is required: -path")
	}
	if OptLevelStr != "" {
		if level, err := log.ParseLevel(OptLevelStr); err != nil {
			return errors.New(fmt.Sprintf("syntex error: %s", OptLevelStr))
		} else {
			OptLevel = level
		}
	}
	return
}

func main() {
	if len(os.Args) == 1 {
		PrintUsageToStderr()
		os.Exit(1)
	}
	if err := Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		PrintUsageToStderr()
		os.Exit(1)
	}

	// start grpc
	l, err := net.Listen("tcp", OptAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to listen: %v\n", err)
		os.Exit(2)
	}

	_s := NewSrv()
	if err := _s.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Srv.Init fail: %v\n", err)
		os.Exit(2)
	}
	defer _s.Close()

	s := grpc.NewServer()
	proto.RegisterLogServer(s, _s)
	go s.Serve(l)
	defer s.GracefulStop()

	quit := make(chan os.Signal)
	signal.Ignore(syscall.SIGPIPE)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
