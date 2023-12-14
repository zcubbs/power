package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zcubbs/power/cmd/server/config"
	"github.com/zcubbs/power/cmd/server/db"
	pb "github.com/zcubbs/power/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
)

type Server struct {
	pb.UnimplementedBlueprintServiceServer

	store db.Store
	cfg   config.Configuration
}

func NewServer(store db.Store, cfg config.Configuration) (*Server, error) {
	return &Server{
		store: store,
		cfg:   cfg,
	}, nil
}

func (s *Server) StartGrpcServer() {
	grpcLogger := grpc.UnaryInterceptor(GrpcLogger)

	var tlsOpt grpc.ServerOption
	if s.cfg.GrpcServer.Tls.Enabled {
		var err error
		tlsOpt, err = newServerTlsOptions(s.cfg.GrpcServer)
		if err != nil {
			log.Fatal("cannot create new server tls options", "error", err)
		}
	}

	grpcServer := grpc.NewServer(grpcLogger, tlsOpt)
	pb.RegisterBlueprintServiceServer(grpcServer, s)

	if s.cfg.GrpcServer.EnableReflection {
		reflection.Register(grpcServer)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcServer.Port))
	if err != nil {
		log.Fatal("cannot listen", "error", err, "port", s.cfg.GrpcServer.Port)
	}

	log.Info("🟢 starting grpc server", "port", s.cfg.GrpcServer.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("cannot start grpc server", "error", err)
	}
}

func (s *Server) StartHttpGateway() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := newGrpcRuntimeServerMux()

	err := pb.RegisterBlueprintServiceHandlerServer(ctx, grpcMux, s)
	if err != nil {
		log.Fatal("cannot register handler server", "error", err)
	}

	mux := http.NewServeMux()
	// add grpc handler
	mux.Handle("/", grpcMux)
	handler := HttpLogger(mux)

	httpSrv := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.cfg.HttpServer.Port),
		ReadHeaderTimeout: s.cfg.HttpServer.ReadHeaderTimeout,
		Handler:           handler,
	}

	log.Info("🟢 starting HTTP Gateway server", "port", s.cfg.HttpServer.Port)
	if err := httpSrv.ListenAndServe(); err != nil {
		log.Fatal("cannot start http server", "error", err)
	}
}

func newGrpcRuntimeServerMux() *runtime.ServeMux {
	jsonOpts := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	return runtime.NewServeMux(jsonOpts)
}

func newServerTlsOptions(cfg config.GrpcServerConfig) (grpc.ServerOption, error) {
	// Load the certificates from disk
	certificate, err := tls.LoadX509KeyPair(cfg.Tls.Cert, cfg.Tls.Key)
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %w", err)
	}

	// Create the TLS credentials
	return grpc.Creds(credentials.NewServerTLSFromCert(&certificate)), nil
}
