package config

import "time"

type Configuration struct {
	Debug             bool             `mapstructure:"debug"`
	HttpServer        HttpServerConfig `mapstructure:"http_server"`
	GrpcServer        GrpcServerConfig `mapstructure:"grpc_server"`
	Database          DatabaseConfig   `mapstructure:"database"`
	InitAdminPassword string           `mapstructure:"init_admin_password"`
	S3                S3Config         `mapstructure:"s3"`

	// Version is the version of the application.
	Version string `json:"version"`
	// Commit is the git commit of the application.
	Commit string `json:"commit"`
	// Date is the build date of the application.
	Date string `json:"date"`
}

type HttpServerConfig struct {
	Port         int    `mapstructure:"port"`
	AllowOrigins string `mapstructure:"allow_origins"`
	AllowHeaders string `mapstructure:"allow_headers"`
	TZ           string `mapstructure:"tz"`
	// ReadHeaderTimeout is the amount of time allowed to read request headers. Default values: '3s'
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout"`
}

type GrpcServerConfig struct {
	Port             int       `mapstructure:"port"`
	EnableReflection bool      `mapstructure:"enable_reflection"`
	Tls              TlsConfig `mapstructure:"tls"`
}

type TlsConfig struct {
	Enabled bool   `mapstructure:"enabled"`
	Cert    string `mapstructure:"cert"`
	Key     string `mapstructure:"key"`
}

type DatabaseType string

const (
	Postgres DatabaseType = "postgres"
)

type DatabaseConfig struct {
	AutoMigration bool           `mapstructure:"auto_migration" json:"auto_migration"`
	Postgres      PostgresConfig `mapstructure:"postgres" json:"postgres"`
}

func (dc *DatabaseConfig) GetDatabaseType() DatabaseType {
	if dc.Postgres.Enabled {
		return Postgres
	}
	return ""
}

type PostgresConfig struct {
	Enabled  bool   `mapstructure:"enabled" json:"enabled"`
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Username string `mapstructure:"username" json:"username"`
	Password string `mapstructure:"password" json:"password"`
	DbName   string `mapstructure:"db_name" json:"db_name"`
	SslMode  bool   `mapstructure:"ssl_mode" json:"ssl_mode"`
	Verbose  bool   `mapstructure:"verbose" json:"verbose"`
	CertPem  string `mapstructure:"cert_pem"`
	CertKey  string `mapstructure:"cert_key"`
	// MaxConns is the maximum number of connections in the pool. Default value: 10
	MaxConns int32 `mapstructure:"max_conns" json:"max_conns"`
	// MinConns is the minimum number of connections in the pool. Default value: 2
	MinConns int32 `mapstructure:"min_conns" json:"min_conns"`
}

type S3Config struct {
	Endpoint       string `mapstructure:"endpoint" json:"endpoint"`
	AccessKey      string `mapstructure:"access_key" json:"access_key"`
	SecretKey      string `mapstructure:"secret_key" json:"secret_key"`
	UseSSL         bool   `mapstructure:"use_ssl" json:"use_ssl"`
	BucketName     string `mapstructure:"bucket_name" json:"bucket_name"`
	DownloadHost   string `mapstructure:"download_host" json:"download_host"`
	DownloadScheme string `mapstructure:"download_scheme" json:"download_scheme"`
}
