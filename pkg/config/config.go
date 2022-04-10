package config

import (
	"strings"

	"github.com/spf13/viper"
)

// Config configuration contract
type Config interface {
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	initialize(basepath, configPath string)
}

type viperConfig struct{}

func (v *viperConfig) initialize(basepath, configPath string) {
	viper.SetEnvPrefix("amartha_beyond_config")
	viper.AutomaticEnv()

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigType("json")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetString get string value from config file.
func (v *viperConfig) GetString(key string) string {
	return viper.GetString(key)
}

// GetInt get Int value from config file.
func (v *viperConfig) GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool get boolean value from config file.
func (v *viperConfig) GetBool(key string) bool {
	return viper.GetBool(key)
}

// NewViperConfig new instance of configuration
func NewViperConfig(basepath, configPath string) Config {
	v := &viperConfig{}
	v.initialize(basepath, configPath)

	return v
}
