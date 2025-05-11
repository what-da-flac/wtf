package loggers

import grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"

func ZapOptionIgnoreHealthCheck() grpc_zap.Option {
	const endpointName = "/grpc.health.v1.Health/Check"
	return grpc_zap.WithDecider(func(fullMethodName string, err error) bool {
		// will not log gRPC calls if it was a call to healthcheck and no error was raised
		if err == nil && fullMethodName == endpointName {
			return false
		}
		// by default any other endpoint will be logged
		return true
	})
}
