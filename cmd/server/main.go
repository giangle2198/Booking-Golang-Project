package main

import (
	"booking-center-service/config"
	"booking-center-service/internal/registry"
	"booking-center-service/internal/util"
	"booking-center-service/pb"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	_ "booking-center-service/cmd/server/statik"
	bimClientPB "booking-identity-management/pb"
	"booking-libs/log"

	"github.com/rs/cors"
)

func main() {
	registry.BuildDIContainer()

	cfg := registry.GetDependency("Config").(*config.Config)
	err := log.InitZap(cfg.App, cfg.Environment, cfg.SensitiveFields)
	if err != nil {
		zap.S().Panic("Can't init zap logger", zap.Error(err))
	}
	zap.S().Infof("Init Main API")

	api := registry.GetDependency("API").(pb.APIServer)

	var server *grpc.Server
	var gateway *http.Server

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Base.GRPCAddress))
		if err != nil {
			zap.S().Panic("Server: Can't create listener", zap.Error(err))
		}

		zapFunc := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool {
			return true
		}

		bimClient := registry.GetDependency("bimClientAdapter").(bimClientPB.BIMAPIClient)

		server = grpc.NewServer(
			grpc_middleware.WithUnaryServerChain(
				util.JWTAuthentication(bimClient, cfg.ExcludedPaths),
				grpc_zap.PayloadUnaryServerInterceptor(zap.S().Desugar(), zapFunc),
			),
		)
		pb.RegisterAPIServer(server, api)

		if err := server.Serve(listener); err != nil {
			zap.S().Panic("Server: Failed to serve", zap.Error(err))
		}
	}()

	zap.S().Infof("Server is started on port %d", cfg.Base.GRPCAddress)

	go func() {
		statik, err := fs.New()
		if err != nil {
			zap.S().Panic("Can't create statik", zap.Error(err))
		}
		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
		opts := []grpc.DialOption{grpc.WithInsecure()}

		err = pb.RegisterAPIHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", cfg.Base.GRPCAddress), opts)
		if err != nil {
			return
		}

		httpMux := http.NewServeMux()
		httpMux.Handle("/bcs/swagger-ui/", http.StripPrefix("/bcs/swagger-ui/", http.FileServer(statik)))
		httpMux.Handle("/", mux)

		var httpHandler http.Handler = httpMux
		if cfg.EnabledCors {
			corsHandler := cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowCredentials: true,
				AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
				AllowedHeaders:   []string{"*"},
				ExposedHeaders: []string{"Grpc-Metadata-Authorization",
					"Content-Type", "Content-Disposition",
					"File-Name",
					"Content-Transfer-Encoding",
					"Grpc-Metadata-Custom-Header-Additional-Info", "Authorization"},
				Debug: true,
			})
			httpHandler = corsHandler.Handler(httpMux)
		}

		gateway = &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Base.HTTPAddress),
			Handler: httpHandler,
		}

		if err := gateway.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Panic("Gateway: Failed to listen and serve", zap.Error(err))
		}
	}()

	zap.S().Infof("Gateway is started on port %d", cfg.Base.HTTPAddress)

	zap.S().Info("*****RUNNING******")

	signals := make(chan os.Signal, 1)
	shutdown := make(chan bool, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// Stop http gateway
		if err := gateway.Shutdown(ctx); err != nil {
			zap.S().Fatalw("Failed to shutdown gateway", zap.Error(err))
		}
		// Stop grpc server
		server.GracefulStop()
		shutdown <- true
	}()
	<-shutdown
	zap.S().Info("*****SHUTDOWN*****")
}
