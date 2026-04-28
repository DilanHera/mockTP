package app

import (
	internal "github.com/DilanHera/mockTP/internal"
	"github.com/DilanHera/mockTP/internal/kafka"
	"github.com/DilanHera/mockTP/internal/services"
	"github.com/DilanHera/mockTP/internal/store"
)

type App struct {
	Helper       internal.Helper
	ApiInfoStore store.ApiInfoStore
	Service      services.Service
	Kafka        kafka.Kafka
}

func NewApp(kafkaConfig kafka.KafkaAppConfig) *App {
	db, err := store.Open()
	apiInfoStore := store.NewApiInfoStore(db)
	if err != nil {
		panic(err)
	}
	var kafkaInstance kafka.Kafka
	if kafkaConfig.Broker != "" {
		kafkaInstance = kafka.NewKafka(kafkaConfig)
	}
	return &App{
		Helper:       internal.NewHelper(),
		ApiInfoStore: *apiInfoStore,
		Service:      services.NewService(apiInfoStore),
		Kafka:        kafkaInstance,
	}
}
