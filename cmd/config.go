package main

import (
	"fmt"

	"github.com/ansharw/rest-api/services"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func getConfig() (*services.Config, error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("cockatoo")
	viper.SetConfigType("json")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("[Config]getConfig - failed ReadInConfig() [%s]", err)
	}

	conf := &services.Config{}
	errUnMarshal := viper.Unmarshal(conf)
	if errUnMarshal != nil {
		return conf, fmt.Errorf("[Config]getConfig - failed unmarshal() [%s]", errUnMarshal)
	}

	validate := validator.New()
	err = validate.Struct(conf)
	if err != nil {
		return conf, fmt.Errorf("[Config]getConfig - fail do validation [%s]", err)
	}

	return conf, err
}
