package ifaces

//go:generate moq -out ../mocks/identifier.go -pkg mocks . Identifier
type Identifier interface {
	UUIDv4() string
}
