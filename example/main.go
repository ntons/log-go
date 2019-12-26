package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ntons/log-go"
	_ "github.com/ntons/log-go/logrus"
	_ "github.com/ntons/log-go/zap"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "config.{json,yaml}")
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Fail to read", os.Args, ",", err)
		os.Exit(1)
	}
	switch ext := filepath.Ext(os.Args[1]); ext {
	case ".json":
		err = log.ConfigFromJSON(b)
	case ".yaml":
		err = log.ConfigFromYAML(b)
	default:
		fmt.Println("Bad config ext:", ext)
	}
	if err != nil {
		fmt.Println("Fail to config from", os.Args[1], ",", err)
		return
	}
	log.L("zapExample").Info("zap::log")
	log.L("zapExample").Infof("%s::logf", "zap")
	log.L("zapExample").Infow("zap::logw", log.Fields{"foo": "bar"})
	log.L("zapExample").With(log.Fields{"foo": "bar"}).Info("zap::with")

	log.L("logrusExample").Info("logrus::log")
	log.L("logrusExample").Infof("%s::logf", "logrus")
	log.L("logrusExample").Infow("logrus::logw", log.Fields{"foo": "bar"})
	log.L("logrusExample").With(log.Fields{"foo": "bar"}).Info("logrus::with")
}
