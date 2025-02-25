package setting

import "github.com/prometheus/client_golang/prometheus"

// Struct Config
type Config struct {
	Server   ServerSetting   `mapstructure:"server"`
	MySQL    MySQLSetting    `mapstructure:mysql"`
	Logger   LoggerSetting   `mapstructure:"logger"`
	Redis    RedisSetting    `mapstructure:"redis"`
	SendGrid SendGridSetting `mapstructure:"send_grid"`
	//RabbitMQ RabbitMQSetting `mapstructure:"rabbitmq"`
	Jwt JwtSetting `mapstructure:"jwt"`
}

// SendGrid Struct Settings
type SendGridSetting struct {
	APIKey string `mapstructure:"api_key"`
}

// Server Struct
type ServerSetting struct {
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

// My SQL Struct
type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifetime"`
}

// Redis Struct
type RedisSetting struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

// Logger Struct
type LoggerSetting struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"file_log_name"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

// Jwt struct
type JwtSetting struct {
	TOKEN_HOUR_LIFESPAN uint   `mapstructure:"token_hour_lifespan"`
	JWT_EXPIRATION      string `mapstructure:"jwt_expiration"`
	API_SECRET          string `mapstructure:"api_secret"`
}

// Prometheus setting structure
type PrometheusSetting struct {
	RequestCount      *prometheus.CounterVec
	ErrorCount        *prometheus.CounterVec
	RequestDuration   *prometheus.HistogramVec
	RequestSizeBytes  *prometheus.HistogramVec
	ResponseSizeBytes *prometheus.HistogramVec
}
