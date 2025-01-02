package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:user`
		Password string `mapstructure:password`
		Host     string `mapstructure:host`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Server port: ", viper.Get("server.port"))
	fmt.Println("Security Jwt key: ", viper.GetString("security.jwt.key"))

	// config structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("Config port::", config.Server.Port)

	for _, v := range config.Databases {
		fmt.Println("Config port::", v.User)
	}
}
