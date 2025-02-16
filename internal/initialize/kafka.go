package initialize

import (
	"example.com/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
	"log"
)

// Initial kafka Producer
var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:19092"),
		Topic:    "otp-auth-topic", // topic
		Balancer: &kafka.LeastBytes{},
	}

}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatalf("Failed to close kafka producer: %v", err)
	}
}
