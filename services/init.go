package services

var env *Config

type Config struct {
	Adapter struct {
		Redis struct {
			Host         string `mapstructure:"host" validate:"required"`
			Port         string `mapstructure:"port" validate:"required"`
			DB           int    `mapstructure:"db" validate:""`
			Password     string `mapstructure:"password" validate:""`
			Expire       int64  `mapstructure:"expire" validate:"required"`
			TimeLimit    int64  `mapstructure:"time_limit" validate:"required"`
			RequestLimit int64  `mapstructure:"request_limit" validate:"required"`
		} `mapstructure:"redis"`
		Mysql struct {
			Host                       string `mapstructure:"host" validate:"required"`
			Port                       string `mapstructure:"port" validate:"required"`
			Username                   string `mapstructure:"username" validate:"required"`
			Password                   string `mapstructure:"password" validate:"required"`
			Database                   string `mapstructure:"database" validate:"required"`
			MaxRetryConnect            int    `mapstructure:"max_retry_connect" validate:""`
			MaxOpenConnection          int    `mapstructure:"max_open_connection" validate:""`
			LifetimeDurationConnection int    `mapstructure:"lifetime_duration_connection" validate:""`
		} `mapstructure:"mysql"`
	} `mapstructure:"adapter"`
	Service struct {
		Port               string `mapstructure:"port" validate:"required"`
		TimeoutResponseApi int    `mapstructure:"timeout_response_api" validate:"required"`
	} `mapstructure:"service"`
}

func InitSetEnv(data *Config) {
	env = data
}
