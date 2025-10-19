package configs

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	BrasilApiUrl string `mapstructure:"BRASIL_API_URL"`
	ViaCepApiUrl string `mapstructure:"VIACEP_API_URL"`
}

func LoadConfig(path string) (*Configuration, error) {
	var config *Configuration
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return config, err
}
