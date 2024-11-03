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
	return zl.Sugar()
}

func MustNewProductionLogger() ifaces.Logger {
	zl, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return zl.Sugar()
}
