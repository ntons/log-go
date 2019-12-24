package log

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Engine string          `json:"engine"`
	Config json.RawMessage `json:"config"`
}

func ConfigFromJSON(b []byte) (err error) {
	var root map[string]*Config
	if err = json.Unmarshal(b, &root); err != nil {
		return
	}
	var (
		success = false
		loggers = make(map[string]Logger)
	)
	defer func() {
		if !success {
			for _, logger := range loggers {
				logger.Close()
			}
		}
	}()
	for name, cfg := range root {
		builder, err := newLoggerBuilder(name)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(cfg.Config, builder); err != nil {
			return err
		}
		logger, err := builder.Build()
		if err != nil {
			return err
		}
		loggers[name] = logger
	}
	// replace global loggers
	replaceGlobalLoggers(loggers)
	return
}

func ConfigFromYAML(b []byte) (err error) {
	var cfg interface{}
	if err = yaml.Unmarshal(b, &cfg); err != nil {
		return
	}
	if b, err = json.Marshal(cfg); err != nil {
		return
	}
	return ConfigFromJSON(b)
}
