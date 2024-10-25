package grpcs

import "google.golang.org/protobuf/proto"

func Unmarshal(data []byte, m proto.Message) error {
	return proto.Unmarshal(data, m)
}
