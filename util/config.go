package util

import "github.com/spf13/viper"

// * Config stores all configuration of the applicaton
// * The values are read by viper from a config file or enviroment variables
type Config struct {
	DBDrive       string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADRESS"`
}

// * LoadConfig reads configuration from a file or enviroment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
