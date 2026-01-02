package kafka

import (
	"github.com/cristiano-pacheco/go-starter-kit/internal/shared/modules/config"
	"github.com/cristiano-pacheco/go-starter-kit/pkg/kafka"
)

func NewKafkaFacade(config config.Config) kafka.Builder {
	kafkaConfig := kafka.Config{
		Address: config.Kafka.Address,
	}
	return kafka.MustNewKafkaBuilder(kafkaConfig)
}
