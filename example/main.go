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
	log.L("zapExample").Infow("test zap", log.Fields{"foo": "bar"})
	log.L("logrusExample").Info("test logrus")
	log.L("logrusExample").Infof("test %s", "logrus")
	log.L("logrusExample").Infow("test logrus", log.Fields{"foo": "bar"})
}
