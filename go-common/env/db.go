package env

import "github.com/spf13/viper"

type DB struct {
	URL string
}

const (
	envVarDbUrl = "DB_URL"
)

func newDB() DB {
	return DB{
		URL: viper.GetString(envVarDbUrl),
	}
}
