package config

import (
	"log"

	"github.com/spf13/viper"
	"gitlab.com/auth-service/external/constants"
)

// Repository repository
type Repository interface {
	NewConfig(string) *Config
	GetString(string) string
}

// Config return struct of viper
type Config struct {
	conf *viper.Viper
}

// NewConfig return new instance of config
func NewConfig(configPath string) *Config {
	v := viper.New()
	v.AutomaticEnv()
	v.SetConfigType("yaml")
	v.SetConfigName("application")
	v.AddConfigPath(configPath)
	v.AddConfigPath(constants.EnvConfigPath)
	err := v.ReadInConfig()
	if err != nil {
		log.Println(err.Error())
	}

	return &Config{
		conf: v,
	}
}

// GetString interface of
func (c *Config) GetString(str string) string {
	res := c.conf.GetString(str)
	return res
}
