package api

import (
	"context"
	"crypto/tls"
	"embed"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/zcubbs/power/cmd/server/config"
	"github.com/zcubbs/power/cmd/server/db"
	"github.com/zcubbs/power/pkg/miniohelper"
	pb "github.com/zcubbs/power/proto/gen/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"io/fs"
	"net"
	"net/http"
	"time"
)

type Server struct {
	pb.UnimplementedBlueprintServiceServer

	store     db.Store
	s3Client  *miniohelper.MinIOClient
	cfg       config.Configuration
	embedOpts []EmbedAssetsOpts
}

func NewServer(store db.Store, cfg config.Configuration, embedOpts ...EmbedAssetsOpts) (*Server, error) {
	s := &Server{
		store:     store,
		cfg:       cfg,
		embedOpts: embedOpts,
	}

	// setup s3 client
	setupS3Client(s)

	return s, nil
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
	} else {
		log.Warn("🔴 grpc server is running without TLS")
		tlsOpt = grpc.EmptyServerOption{}
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

	// add embedded assets handler
	for _, opts := range s.embedOpts {
		mux.Handle(opts.Path, newFileServerHandler(opts))
	}

	// add grpc handler
	mux.Handle("/", grpcMux)
	handler := HttpLogger(mux)

	// Cors
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	handler = handlers.CORS(origins, methods, headers)(handler)

	// server options
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

type EmbedAssetsOpts struct {
	// The directory to embed.
	Dir    embed.FS
	Path   string
	Prefix string
}

func newFileServerHandler(opts EmbedAssetsOpts) http.Handler {
	log.Info("serving embedded assets", "path", opts.Path)
	sub, err := fs.Sub(opts.Dir, opts.Prefix)
	if err != nil {
		log.Fatal("cannot serve embedded assets", "error", err)
	}
	dir := http.FileServer(http.FS(sub))

	return http.StripPrefix(opts.Path, dir)
}

func setupS3Client(s *Server) {
	retryCount := 0

	for {
		if retryCount >= s.cfg.S3.ConnectionRetryCount {
			log.Fatal("maximum retries reached for setting up S3 client, exiting")
			return // Or handle this situation appropriately
		}

		var err error
		s.s3Client, err = miniohelper.New(s.cfg.S3.Endpoint, s.cfg.S3.AccessKey, s.cfg.S3.SecretKey, s.cfg.S3.UseSSL)
		if err != nil {
			log.Error("failed to create S3 client, retrying", "error", err, "retryCount", retryCount)
			retryCount++
			time.Sleep(5 * time.Second)
			continue
		}

		// Check if the bucket exists
		bucketExists, err := s.s3Client.BucketExists(s.cfg.S3.BucketName)
		if err != nil {
			log.Error("failed to check if bucket exists, retrying", "error", err, "bucketName", s.cfg.S3.BucketName, "retryCount", retryCount)
			retryCount++
			time.Sleep(5 * time.Second)
			continue
		}

		if !bucketExists {
			// Create the bucket if it does not exist
			err = s.s3Client.MakeBucket(s.cfg.S3.BucketName)
			if err != nil {
				log.Error("failed to create bucket, retrying", "error", err, "bucketName", s.cfg.S3.BucketName, "retryCount", retryCount)
				retryCount++
				time.Sleep(5 * time.Second)
				continue
			}
			log.Info("bucket created", "bucketName", s.cfg.S3.BucketName)
		} else {
			log.Info("bucket already exists", "bucketName", s.cfg.S3.BucketName)
		}

		log.Info("S3 client setup successful")
		break
	}
}
