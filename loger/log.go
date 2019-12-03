package loger

import (
	"os"

	"go.etcd.io/etcd/pkg/logutil"
	"go.uber.org/zap"
)

var Sugar *zap.SugaredLogger

const (
	// DebugMode indicates gin mode is debug.
	DebugMode = "debug"
	// ReleaseMode indicates gin mode is release.
	ReleaseMode = "release"
	// TestMode indicates gin mode is test.
	TestMode = "test"
)

// 给定一个env key或者一个默认值
func SugaredLog(env string, def string) (sugar *zap.SugaredLogger, err error) {
	var (
		level string
		lg    *zap.Logger
	)
	if lg, err = zap.NewProduction(); err != nil {
		return
	}
	defer lg.Sync()
	cfg := logutil.DefaultZapLoggerConfig
	switch env {
	case "":
		level = ReleaseMode
	default:
		level = os.Getenv(env)
	}

	switch level {
	case DebugMode:
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case ReleaseMode:
		cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
	default:
		switch def {
		case DebugMode:
			cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		case ReleaseMode:
			cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		default:
			cfg.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		}
	}

	if lg, err = cfg.Build(); err != nil {
		return
	}
	sugar = lg.Sugar()
	return
}
