package initialize

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/tqlong1609/go_backend_ecommerce/global"
)

func InitConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
