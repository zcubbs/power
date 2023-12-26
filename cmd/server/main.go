package main

import (
	"flag"
	"github.com/charmbracelet/log"
	"github.com/zcubbs/power/cmd/server/api"
	"github.com/zcubbs/power/cmd/server/config"
	"github.com/zcubbs/power/cmd/server/docs"
	"github.com/zcubbs/power/cmd/server/utils"
	"github.com/zcubbs/power/pkg/blueprint"
	"github.com/zcubbs/power/pkg/designer"
	"github.com/zcubbs/power/pkg/plugin"
	"os"
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

	// Set the timezone
	err = os.Setenv("TZ", cfg.HttpServer.TZ)
	if err != nil {
		log.Error("failed to set timezone", "error", err)
	}
	utils.CheckTimeZone()

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
	if cfg.Blueprint.EnableBuiltins {
		err := designer.EnableBuiltinGenerators()
		if err != nil {
			log.Fatal("failed to register builtin blueprints", "error", err)
		}
	}

	// Register Plugins
	if cfg.Blueprint.EnablePlugins {
		if cfg.Blueprint.PluginDir == "" {
			log.Fatal("plugins are enabled but no plugin dir is set")
		}
		err := plugin.LoadPlugins(cfg.Blueprint.PluginDir)
		if err != nil {
			log.Fatal("failed to load plugin", "error", err)
		}
	}

	for _, bp := range blueprint.GetAllBlueprints() {
		log.Info("registered blueprint",
			"id", bp.Spec.ID,
			"name", bp.Spec.Name,
			"version", bp.Spec.Version,
		)
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
