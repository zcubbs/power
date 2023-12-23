package config

const (
	DefaultDbName = "power"
	Localhost     = "127.0.0.1"
	HttpPort      = 8000
	GrpcPort      = 9000
)

var (
	defaults = map[string]interface{}{
		"debug":                           false,
		"http_server.port":                HttpPort,
		"http_server.allow_origins":       "http://localhost:3000",
		"http_server.allow_headers":       "Origin, Content-Type, Accept",
		"http_server.tz":                  "UTC",
		"http_server.enable_print_routes": false,
		"http_server.read_header_timeout": "3s",
		"grpc_server.port":                GrpcPort,
		"grpc_server.enable_reflection":   true,
		"grpc_server.tls.enabled":         false,
		"grpc_server.tls.cert":            "",
		"grpc_server.tls.key":             "",
		"database.auto_migration":         true,
		"database.postgres.enabled":       true,
		"database.postgres.host":          Localhost,
		"database.postgres.port":          5432,
		"database.postgres.username":      "postgres",
		"database.postgres.password":      "postgres",
		"database.postgres.db_name":       DefaultDbName,
		"database.postgres.ssl_mode":      false,
		"database.postgres.verbose":       false,
		"database.postgres.cert_pem":      "",
		"database.postgres.cert_key":      "",
		"database.postgres.max_conns":     10,
		"database.postgres.min_conns":     4,
		"s3.endpoint":                     "localhost:8333",
		"s3.access_key":                   "storage_access_key",
		"s3.secret_key":                   "storage_secret_key",
		"s3.use_ssl":                      false,
		"s3.bucket_name":                  "power",
	}

	envKeys = map[string]string{
		"debug":                           "DEBUG",
		"init_admin_password":             "INIT_ADMIN_PASSWORD",
		"http_server.port":                "HTTP_SERVER_PORT",
		"http_server.allow_origins":       "HTTP_SERVER_ALLOW_ORIGINS",
		"http_server.allow_headers":       "HTTP_SERVER_ALLOW_HEADERS",
		"http_server.tz":                  "HTTP_SERVER_TZ",
		"http_server.enable_print_routes": "HTTP_SERVER_ENABLE_PRINT_ROUTES",
		"http_server.read_header_timeout": "HTTP_SERVER_READ_HEADER_TIMEOUT",
		"grpc_server.port":                "GRPC_SERVER_PORT",
		"grpc_server.enable_reflection":   "GRPC_SERVER_ENABLE_REFLECTION",
		"grpc_server.tls.enabled":         "GRPC_SERVER_TLS_ENABLED",
		"grpc_server.tls.cert":            "GRPC_SERVER_TLS_CERT",
		"grpc_server.tls.key":             "GRPC_SERVER_TLS_KEY",
		"database.auto_migration":         "DATABASE_AUTO_MIGRATION",
		"database.postgres.enabled":       "DATABASE_POSTGRES_ENABLED",
		"database.postgres.host":          "DATABASE_POSTGRES_HOST",
		"database.postgres.port":          "DATABASE_POSTGRES_PORT",
		"database.postgres.username":      "DATABASE_POSTGRES_USERNAME",
		"database.postgres.password":      "DATABASE_POSTGRES_PASSWORD",
		"database.postgres.database":      "DATABASE_POSTGRES_DATABASE",
		"database.postgres.ssl_mode":      "DATABASE_POSTGRES_SSL_MODE",
		"database.postgres.verbose":       "DATABASE_POSTGRES_VERBOSE",
		"database.postgres.cert_pem":      "DATABASE_POSTGRES_CERT_PEM",
		"database.postgres.cert_key":      "DATABASE_POSTGRES_CERT_KEY",
		"database.postgres.max_conns":     "DATABASE_POSTGRES_MAX_CONNS",
		"database.postgres.min_conns":     "DATABASE_POSTGRES_MIN_CONNS",
		"s3.endpoint":                     "S3_ENDPOINT",
		"s3.access_key":                   "S3_ACCESS_KEY",
		"s3.secret_key":                   "S3_SECRET_KEY",
		"s3.use_ssl":                      "S3_USE_SSL",
		"s3.bucket_name":                  "S3_BUCKET_NAME",
	}
)
