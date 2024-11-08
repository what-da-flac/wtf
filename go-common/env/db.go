package env

import "github.com/spf13/viper"

type DB struct {
	EnableMigrations bool
	URL              string
}

const (
	envVarDbUrl              = "DB_URL"
	envVarDbEnableMigrations = "DB_ENABLE_MIGRATIONS"
)

func newDB() DB {
	return DB{
		EnableMigrations: viper.GetBool(envVarDbEnableMigrations),
		URL:              viper.GetString(envVarDbUrl),
	}
}
