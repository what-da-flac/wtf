package imodel

type EndpointPermission struct {
	Path     string   `yaml:"path"`
	Method   string   `yaml:"method"`
	Roles    []string `yaml:"roles"`
	Insecure bool     `yaml:"insecure"`
}

func (x *EndpointPermission) AsEndpointPath() EndpointPath {
	return EndpointPath(x.Path)
}
