package log

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Engine string `json:"engine"`
}

func ConfigFromJSON(b []byte) (err error) {
	var root map[string]json.RawMessage
	if err = json.Unmarshal(b, &root); err != nil {
		return
	}
	var loggers = make(map[string]Logger)
	for name, raw := range root {
		var (
			cfg     Config
			builder Builder
			logger  Logger
		)
		if err := json.Unmarshal(raw, &cfg); err != nil {
			break
		}
		if builder, err = newLoggerBuilder(cfg.Engine); err != nil {
			break
		}
		if err = json.Unmarshal(raw, builder); err != nil {
			break
		}
		if logger, err = builder.Build(); err != nil {
			break
		}
		loggers[name] = logger
	}
	if err != nil {
		for _, logger := range loggers {
			logger.Close()
		}
		return
	}
	// replace global loggers
	replaceGlobalLoggers(loggers)
	return
}

func YAMLMap2JSONMap(x interface{}) interface{} {
	switch x := x.(type) {
	case map[interface{}]interface{}:
		m2 := make(map[string]interface{})
		for k, v := range x {
			m2[k.(string)] = YAMLMap2JSONMap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = YAMLMap2JSONMap(v)
		}
		return x
	default:
		return x
	}
}

func YAML2JSON(b []byte) ([]byte, error) {
	var m interface{}
	if err := yaml.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return json.Marshal(YAMLMap2JSONMap(m))
}

func ConfigFromYAML(b []byte) (err error) {
	if b, err = YAML2JSON(b); err != nil {
		return
	}
	return ConfigFromJSON(b)
}
