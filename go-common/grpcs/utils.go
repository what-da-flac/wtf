package grpcs

import "google.golang.org/protobuf/proto"

// Clone returns a deep copy of m.
func Clone(m proto.Message) proto.Message {
	return proto.Clone(m)
}
