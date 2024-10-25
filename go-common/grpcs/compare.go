package grpcs

import (
	"google.golang.org/protobuf/proto"
)

// CompareProtobuf returns true if the JSON representation of both protobuf messages are identical.
func CompareProtobuf(m1, m2 proto.Message) bool {
	d1, err := ToJSON(m1, true, true)
	if err != nil {
		return false
	}
	d2, err := ToJSON(m2, true, true)
	if err != nil {
		return false
	}
	return string(d1) == string(d2)
}
