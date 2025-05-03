package env

import (
	"github.com/spf13/viper"
)

type Path struct {
	Storage string
	Temp    string
}

func NewPathsFromEnvironment() *Path {
	p := &Path{
		Storage: viper.GetString("PATH_STORAGE"),
		Temp:    viper.GetString("PATH_TEMP"),
	}
	return p
}
