package config

var (
	DefaultZapJsonConfig = &Config{
		Zap: &zap.Config{
			Level:            zap.NewAtomicLevel(),
			Encoding:         "json",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		},
	}
	DefaultZapConsoleConfig = &Config{
		Zap: &zap.Config{
			Level:            zap.NewAtomicLevel(),
			Encoding:         "console",
			EncoderConfig:    zap.NewProductionEncoderConfig(),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		},
	}
)
