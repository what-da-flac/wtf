package grpcs

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// MustClientConn returns a grpc connection and panic if errors happen in the process.
func MustClientConn(host string) *grpc.ClientConn {
	zap.S().Infof("trying to connect to grpc host: %s", host)
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return conn
}
