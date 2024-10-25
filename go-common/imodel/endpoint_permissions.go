package imodel

import (
	"fmt"
	"io"
	"strings"

	"gopkg.in/yaml.v3"
)

type EndpointPermissions []*EndpointPermission

func (x EndpointPermissions) Filter(filter func(ep *EndpointPermission) bool) EndpointPermissions {
	var res EndpointPermissions
	for _, v := range x {
		if filter(v) {
			res = append(res, v)
		}
	}
	return res
}

func (x EndpointPermissions) Len() int { return len(x) }

func (x EndpointPermissions) IsSecured(path, method string) bool {
	// filter endpoints based on path/method provided
	filtered := x.Filter(func(ep *EndpointPermission) bool {
		return ep.AsEndpointPath().Equals(path) && strings.EqualFold(ep.Method, method)
	})
	// if no record matches the incoming path/method, that means it is unknown endpoint
	// we should mark as not secured (javascript, images or any other resource)
	if filtered.Len() == 0 {
		return false
	}
	insecure := filtered.Filter(func(ep *EndpointPermission) bool {
		// filter the above with insecure access
		return ep.Insecure
	})
	// consider the use case when either endpoint is marked as insecure
	return insecure.Len() == 0
}

func (x EndpointPermissions) Allow(path, method string, roles ...string) bool {
	roleKeys := make(map[string]struct{})
	for _, role := range roles {
		roleKeys[role] = struct{}{}
	}
	// filter endpoints based on path/method provided
	filtered := x.Filter(func(ep *EndpointPermission) bool {
		return ep.AsEndpointPath().Equals(path) && strings.EqualFold(ep.Method, method)
	})
	withRoles := filtered.Filter(func(ep *EndpointPermission) bool {
		// check at least one role has membership
		for _, role := range ep.Roles {
			if _, ok := roleKeys[role]; ok {
				return true
			}
		}
		return false
	})
	return withRoles.Len() != 0
}

func ParsePermissions(r io.ReadCloser) (EndpointPermissions, error) {
	var res EndpointPermissions
	if r == nil {
		return nil, fmt.Errorf("missing parameter: r")
	}
	defer func() { _ = r.Close() }()
	if err := yaml.NewDecoder(r).Decode(&res); err != nil {
		return nil, err
	}
	return res, nil
}
