package log

import "go.uber.org/zap"

var Log *zap.Logger

func SetLogger() {
	logger, _ := zap.NewProduction(zap.WithCaller(false))
	defer logger.Sync()

	Log = logger
}
