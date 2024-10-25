package environment

import "github.com/spf13/viper"

type Sentry struct {
	DSN string
}

const (
	envVarSentryDSN = "SENTRY_DSN"
)

func newSentry() Sentry {
	return Sentry{
		DSN: viper.GetString(envVarSentryDSN),
	}
}
