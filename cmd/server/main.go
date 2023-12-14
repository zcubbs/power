package main

import (
	"flag"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/cmd/server/api"
	"github.com/zcubbs/power/cmd/server/config"
)

var (
	Version = "dev"
	Commit  = "none"
	Date    = "unknown"
)

var cfg *config.Configuration

var configPath = flag.String("config", "", "Path to the configuration file")

func init() {
	flag.Parse()

	// Load configuration
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		log.Fatalf("Error loading configuration error=%s", err)
	}

	cfg.Version = Version
	cfg.Commit = Commit
	cfg.Date = Date
}

func main() {
	// Init context
	// ctx := context.Background()

	// Database Migration
	// TODO: implement migration.Run(cfg.Database)

	// Connect to database
	// TODO: implement dbUtil.Connect(ctx, cfg.Database)

	// Initialize store
	// TODO: implement db.NewSQLStore(conn)

	// Initialize admin user
	// TODO: implement dbUtil.InitAdminUser(store, cfg)

	// Create gRPC Server
	gs, err := api.NewServer(nil, *cfg)
	if err != nil {
		log.Fatal("failed to create grpc server", "error", err)
	}

	// Start HTTP Gateway
	go gs.StartHttpGateway()

	// Start gRPC Server
	gs.StartGrpcServer()
}
