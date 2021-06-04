package config

import (
	"go.uber.org/zap"

	log "github.com/ntons/log-go"
	myzap "github.com/ntons/log-go/zap"
)

type Config struct {
	Level log.Level   `json:"level" yaml:"level"`
	Zap   *zap.Config `json:"zap" yaml:"zap"`
}

func (cfg *Config) Use() error {
	if cfg.Zap != nil {
		zlogger, err := cfg.Zap.Build()
		if err != nil {
			return err
		}
		log.SetStdLogger(myzap.New(zlogger, cfg.Zap.Level))
	}
	if cfg.Level != 0 {
		log.SetStdLevel(cfg.Level)
	}
	return nil
}
