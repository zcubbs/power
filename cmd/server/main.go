package main

import (
	"flag"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/blueprint"
	"github.com/zcubbs/power/cmd/server/api"
	"github.com/zcubbs/power/cmd/server/config"
	"github.com/zcubbs/power/cmd/server/docs"
	"github.com/zcubbs/power/designer"
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
	log.Info("loading configuration...")
	var err error
	cfg, err = config.Load(*configPath)
	if err != nil {
		log.Fatal("failed to load configuration", "error", err)
	}

	cfg.Version = Version
	cfg.Commit = Commit
	cfg.Date = Date

	if cfg.Debug {
		log.SetLevel(log.DebugLevel)
		config.PrintConfiguration(*cfg)
	}

	log.Info("loaded configuration")
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

	// Register Builtin Blueprints
	err := designer.EnableBuiltinGenerators()
	if err != nil {
		log.Fatal("failed to register builtin blueprints", "error", err)
	}

	for k, bp := range blueprint.GetAllBlueprintSpecs() {
		log.Info("registered blueprint", "id", k, "name", bp.Name)
	}

	// Create gRPC Server
	gs, err := api.NewServer(nil, *cfg, api.EmbedAssetsOpts{
		Dir:    docs.SwaggerDist,
		Path:   "/swagger/",
		Prefix: "swagger",
	})
	if err != nil {
		log.Fatal("failed to create grpc server", "error", err)
	}

	// Start HTTP Gateway
	go gs.StartHttpGateway()

	// Start gRPC Server
	gs.StartGrpcServer()
}
