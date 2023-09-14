package config_swagger //nolint:all

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("models.output", "")
	viper.SetDefault("models.specpath", "")
	viper.SetDefault("server.output", "")
	viper.SetDefault("server.specpath", "")
}
