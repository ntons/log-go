package config_test

import (
	"encoding/json"
	"testing"

	"github.com/ntons/log-go"
	"github.com/ntons/log-go/config"
)

func TestConfig(t *testing.T) {
	config.DefaultZapJsonConfig.Use()
	log.Infow("hello", "foo", "bar")

	config.DefaultZapConsoleConfig.Use()
	log.Infow("hello", "foo", "bar")

	b := []byte(`
{
	"level": "DEBUG",
	"zap": {
		"level": "warn",
		"encoding": "json",
		"encoderConfig": {
		    "timeKey":     "tm",
			"timeEncoder": "rfc3339"
		},
		"outputPaths": ["stdout"],
		"errorOutputPaths": ["stderr"]
	}
}`)

	cfg := &config.Config{}
	json.Unmarshal(b, cfg)
	cfg.Use()
	log.Infow("hello", "foo", "bar")
}
