package log

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// [REQUIRED] enabled logger name list
	Loggers []string `json:"loggers" yaml:"loggers"`
	// [OPTIONAL] default logger name, default value: Loggers[0]
	Default string `json:"default" yaml:"default"`
}

type LoggerConfig struct {
	Engine string `json:"engine"`
}

func ConfigFromJSON(b []byte) (err error) {
	var cfg Config
	if err = json.Unmarshal(b, &cfg); err != nil {
		return
	}
	fmt.Println(cfg)
	if len(cfg.Loggers) == 0 {
		return fmt.Errorf("config field %q is required", "loggers")
	}
	if cfg.Default == "" {
		cfg.Default = cfg.Loggers[0]
	}
	var root map[string]json.RawMessage
	if err = json.Unmarshal(b, &root); err != nil {
		return
	}
	var m = make(map[string]Logger)
	for _, name := range cfg.Loggers {
		rawjson, ok := root[name]
		if !ok {
			return fmt.Errorf("logger not configured for %q", name)
		}
		var (
			c LoggerConfig
			b Builder
			l Logger
		)
		if err = json.Unmarshal(rawjson, &c); err != nil {
			break
		}
		if b, err = newLoggerBuilder(c.Engine); err != nil {
			break
		}
		if err = json.Unmarshal(rawjson, b); err != nil {
			break
		}
		if l, err = b.Build(); err != nil {
			break
		}
		m[name] = l
	}
	if err != nil {
		for _, l := range m {
			l.Close()
		}
		return
	}
	if _, ok := m[cfg.Default]; !ok {
		for _, l := range m {
			l.Close()
		}
		return fmt.Errorf("logger not configured for %q", cfg.Default)
	}
	// replace global loggers
	ReplaceLoggers(m, cfg.Default)
	return
}

func ConfigFromYAML(b []byte) (err error) {
	if b, err = yaml2json(b); err != nil {
		return
	}
	return ConfigFromJSON(b)
}
func yamlmap2jsonmap(x interface{}) interface{} {
	switch x := x.(type) {
	case map[interface{}]interface{}:
		m2 := make(map[string]interface{})
		for k, v := range x {
			m2[k.(string)] = yamlmap2jsonmap(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = yamlmap2jsonmap(v)
		}
		return x
	default:
		return x
	}
}
func yaml2json(b []byte) ([]byte, error) {
	var m interface{}
	if err := yaml.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return json.Marshal(yamlmap2jsonmap(m))
}
