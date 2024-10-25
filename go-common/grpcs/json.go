package grpcs

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ToJSON(m proto.Message, useCamelCase, includeNullValues bool) ([]byte, error) {
	marshalOptions := protojson.MarshalOptions{
		Multiline:         false,
		Indent:            "",
		AllowPartial:      false,
		UseProtoNames:     !useCamelCase,
		UseEnumNumbers:    false,
		EmitUnpopulated:   includeNullValues,
		EmitDefaultValues: false,
		Resolver:          nil,
	}
	return marshalOptions.Marshal(m)
}

func FromJSON(data []byte, m proto.Message) error {
	return protojson.Unmarshal(data, m)
}
