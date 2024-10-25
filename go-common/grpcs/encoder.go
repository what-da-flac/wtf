package grpcs

import (
	"github.com/what-da-flac/wtf/go-common/encoders"
	"google.golang.org/protobuf/proto"
)

func Marshal(message proto.Message) ([]byte, error) {
	return proto.Marshal(message)
}

// ProtobufDecoder converts an incoming base64 string into a protobuf structure.
func ProtobufDecoder(encoded64Data string, m proto.Message) error {
	decodedBytes, err := encoders.Decode64(encoded64Data)
	if err != nil {
		return err
	}
	return proto.Unmarshal(decodedBytes, m)
}

// ProtobufEncoder converts protobuf structure into encoded base64 payload.
func ProtobufEncoder(m proto.Message) (string, error) {
	data, err := proto.Marshal(m)
	if err != nil {
		return "", err
	}
	return encoders.Encoder64(data), nil
}
