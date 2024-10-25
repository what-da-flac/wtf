package ifaces

type EndpointDecider interface {
	// Allow returns true if based on path, method and roles provided, the user has access granted to move on.
	Allow(path, method string, roles ...string) bool

	// IsSecured returns true if the endpoint has been marked as insecure, in which case there is no need to process
	// additional authentication checks.
	IsSecured(path, method string) bool
}
