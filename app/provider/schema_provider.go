package provider

import (
	"context"
	"github.com/dns2012/dealls-dating-service/config"
	schemav1 "github.com/dns2012/dealls-dating-service/proto/schema/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterSchemaV1ServerProvider() {
	schemav1.RegisterAuthSchemaServer(config.GrpcServer, SchemaV1UsecaseProvider())
	schemav1.RegisterUserSchemaServer(config.GrpcServer, SchemaV1UsecaseProvider())
	schemav1.RegisterPackageSchemaServer(config.GrpcServer, SchemaV1UsecaseProvider())
}

func RegisterSchemaV1HandlerProvider(ctx context.Context, mux *runtime.ServeMux) {
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := schemav1.RegisterAuthSchemaHandlerFromEndpoint(ctx, mux, *config.GrpcServerEndpoint, dialOptions)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = schemav1.RegisterUserSchemaHandlerFromEndpoint(ctx, mux, *config.GrpcServerEndpoint, dialOptions)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	err = schemav1.RegisterPackageSchemaHandlerFromEndpoint(ctx, mux, *config.GrpcServerEndpoint, dialOptions)
	if err != nil {
		logger.Fatalf(err.Error())
	}
}