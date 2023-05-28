package util

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// reads configuration from file or environment variables.
func LoadConfig(path string, name string) (config Config, err error) {
	var configName string
	viper.AddConfigPath(path)
	if name != "" {
		configName = name
	} else {
		configName = "app"
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
