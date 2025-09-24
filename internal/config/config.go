package config

import "github.com/gomeal/config/pkg/config"

const (
	GrpcServerPort              = config.Key("grpc_server_port")
	ConfigLibDescription        = config.Key("config_lib_description")
	WeatherFetchingCronDuration = config.Key("weather_fetching_cron_duration")
	WeatherFetchingCronEnabled  = config.Key("weather_fetching_cron_enabled")
)

const (
	PostgresPassword = config.Secret("POSTGRES_PASSWORD")
	RabbitMQPort     = config.Secret("RABBITMQ_PORT")
)
