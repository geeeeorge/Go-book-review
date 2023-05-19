package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var conf *Config = nil

// Config represents config of this service
type Config struct {
	appPort   int
	appHost   string
	RDBConfig *RDBConfig
}

// RDBConfig represents config of relational database
type RDBConfig struct {
	port     int
	host     string
	user     string
	password string
	database string
}

// Load loads config
func Load() Config {
	if conf != nil {
		return *conf
	}

	viper.AutomaticEnv()
	c := Config{
		appPort: viper.GetInt("APP_PORT"),
		appHost: viper.GetString("APP_HOST"),
		RDBConfig: &RDBConfig{
			port:     viper.GetInt("DB_PORT"),
			host:     viper.GetString("DB_HOST"),
			user:     viper.GetString("DB_USER"),
			password: viper.GetString("DB_PASSWORD"),
			database: viper.GetString("DB_NAME"),
		},
	}

	return c
}

// AppPort return port number of this service
func (c *Config) AppPort() int {
	return c.appPort
}

// AppHost return app host of this service
func (c *Config) AppHost() string {
	return c.appHost
}

// ConnectionString return string for connecting database
func (r *RDBConfig) ConnectionString() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		r.user, r.password, r.host, r.port, r.database,
	)
}
