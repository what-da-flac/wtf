package imodel

import "strings"

const (
	pathVarOpen  = "{"
	pathVarClose = "}"
)

type EndpointPath string

func (x EndpointPath) String() string {
	return string(x)
}

func (x EndpointPath) isDynamicPart() bool {
	return strings.HasPrefix(x.String(), pathVarOpen) &&
		strings.HasSuffix(x.String(), pathVarClose)
}

func (x EndpointPath) Split() []EndpointPath {
	var result []EndpointPath
	const sep = "/"
	values := strings.Split(x.String(), sep)
	for _, v := range values {
		if v != "" {
			result = append(result, EndpointPath(v))
		}
	}
	return result
}

// Equals compares two paths for equity considering dynamic parts.
func (x EndpointPath) Equals(v string) bool {
	paths := x.Split()
	ep := EndpointPath(v)
	parts := ep.Split()
	if expected, actual := len(paths), len(parts); expected != actual {
		return false
	}
	for i, expected := range paths {
		actual := parts[i]
		if expected.isDynamicPart() {
			// nothing to check here, received path part can be anything
			continue
		}
		if !strings.EqualFold(expected.String(), actual.String()) {
			return false
		}
	}
	return true
}
