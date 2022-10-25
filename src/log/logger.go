package log

import "go.uber.org/zap"

var Log *zap.Logger

func SetLogger() {
	log, _ := zap.NewProduction(zap.WithCaller(false))
	defer func() {
		_ = log.Sync()
	}()
	Log = log
}
