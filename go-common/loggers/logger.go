package loggers

import (
	"github.com/what-da-flac/wtf/go-common/ifaces"
	"go.uber.org/zap"
)

func MustNewDevelopmentLogger() ifaces.Logger {
	zl, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	level := zl.Level()
	if err := level.Set("INFO"); err != nil {
		panic(err)
	}
	return zl.Sugar()
}

func MustNewProductionLogger(logLevel string) ifaces.Logger {
	zl, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	level := zl.Level()
	if err := level.Set(logLevel); err != nil {
		panic(err)
	}
	return zl.Sugar()
}
