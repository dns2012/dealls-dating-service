package main

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/provider"
	"github.com/dns2012/dealls-dating-service/config"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard))

	// NOTE: Register Schema V1 API To GRPC Server
	provider.RegisterSchemaV1ServerProvider()

	config.RunGrpcServer()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// NOTE: Register Schema V1 API Handler (Controller) To GRPC Server And GRPC Gateway
	provider.RegisterSchemaV1HandlerProvider(ctx, config.Mux)

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: config.Mux,
	}

	httpServer.ListenAndServe()

}
