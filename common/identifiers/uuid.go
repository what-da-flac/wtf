package identifiers

import "github.com/google/uuid"

type Identifier struct{}

func NewIdentifier() *Identifier {
	return &Identifier{}
}

func (x Identifier) UUIDv4() string {
	return uuid.NewString()
}
