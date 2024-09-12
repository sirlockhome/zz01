package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server    ServerConfig
	SQLServer SQLServerConfig
	Auth      AuthConfig
}

type ServerConfig struct {
	Port int
	Addr string
}

type SQLServerConfig struct {
	Server          string
	Username        string
	Password        string
	Port            int
	Database        string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifeTime int
}

type RedisConfig struct {
	Addr     string
	Password string
	Username string
	DB       int
}

type AuthConfig struct {
	AccessKey         string
	RefreshKey        string
	AccessExp         int
	RefreshExp        int
	AutoLogoffMinutes int
}

func LoadConfig(filename string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.SetConfigType("yml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileAlreadyExistsError); ok {
			return nil, err
		}
	}

	c, err := parseConfig(v)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func parseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
