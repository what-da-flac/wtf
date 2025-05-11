package env

import (
	"github.com/spf13/viper"
	"github.com/what-da-flac/wtf/openapi/gen/golang"
)

type Path struct {
	keys    map[golang.PathName]string
	Storage string
	Temp    string
}

func NewPathsFromEnvironment() *Path {
	p := &Path{
		Storage: viper.GetString("PATH_STORAGE"),
		Temp:    viper.GetString("PATH_TEMP"),
	}
	p.keys = map[golang.PathName]string{
		golang.PathNameStore: p.Storage,
		golang.PathNameTemp:  p.Temp,
	}
	return p
}

func (x *Path) Resolve(pathName golang.PathName) string {
	if v, ok := x.keys[pathName]; ok {
		return v
	}
	return ""
}
