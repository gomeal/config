package main

import (
	"fmt"

	cfg "github.com/gomeal/config/internal/config"
	"github.com/gomeal/config/pkg/config"
)

func main() {
	configProvider := config.NewProvider(".cfg/local.yaml")

	var (
		grpcServerPort              = configProvider.GetConfigClient().GetValue(cfg.GrpcServerPort).Int()
		configLibDescription        = configProvider.GetConfigClient().GetValue(cfg.ConfigLibDescription).String()
		weatherFetchingCronDuration = configProvider.GetConfigClient().GetValue(cfg.WeatherFetchingCronDuration).Duration()
		weatherFetchingCronEnabled  = configProvider.GetConfigClient().GetValue(cfg.WeatherFetchingCronEnabled).Bool()
	)

	fmt.Printf("grpc_server_port:\n\t%v\n", grpcServerPort)
	fmt.Printf("config_lib_description:\n\t%v\n", configLibDescription)
	fmt.Printf("weather_fetching_cron_duration:\n\t%vs\n", weatherFetchingCronDuration.Seconds())
	fmt.Printf("weather_fetching_cron_enabled:\n\t%v\n", weatherFetchingCronEnabled)

	var (
		postgresPassword = configProvider.GetSecretClient().GetSecret(cfg.PostgresPassword).String()
		rabbitMQPort     = configProvider.GetSecretClient().GetSecret(cfg.RabbitMQPort).Int()
	)

	fmt.Printf("POSTGRES_PASSWORD:\n\t%v\n", postgresPassword)
	fmt.Printf("RABBITMQ_PORT:\n\t%v\n", rabbitMQPort)
}
