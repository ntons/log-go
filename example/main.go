package main

import (
	"fmt"

	"github.com/ntons/log-go"
	_ "github.com/ntons/log-go/logrus"
	_ "github.com/ntons/log-go/zap"
)

const cfg = `{
  "logrusExample": {
    "engine": "logrus",
	"config": {
	  "level": "debug",
	  "reportCaller": true,
	  "out": {
	    "type": "stderr"
	  },
	  "formatter": {
	    "type": "json",
	    "timestampFormat": "2006-01-02T15:04:05.000"
	  }
	}
  },
  "zapExample": {
    "engine": "zap",
	"config": {
      "level": "debug",
      "encoding": "json",
      "outputPaths": ["stdout", "/tmp/logs"],
      "errorOutputPaths": ["stderr"],
      "encoderConfig": {
        "messageKey": "msg",
        "levelKey": "level",
	    "timeKey": "time",
		"callerKey": "caller",
        "levelEncoder": "lowercase",
	    "timeEncoder": "iso8601",
		"callerEncoder": "short"
	  }
    }
  }
}`

func main() {
	if err := log.ConfigFromJSON([]byte(cfg)); err != nil {
		fmt.Println("FromJSON fail: ", err)
		return
	}
	//log.L("zap").Infow("test zap", log.Fields{"foo": "bar"})
	log.L("logrus").Info("test logrus")
	log.L("logrus").Infof("test %s", "logrus")
	log.L("logrus").Infow("test logrus", log.Fields{"foo": "bar"})
}
