package global

import (
	"database/sql"
	"example.com/go-ecommerce-backend-api/pkg/logger"
	"example.com/go-ecommerce-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Mdb           *gorm.DB
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
	//RabbitMQProducer_LOGGERDISCORD *amqp.Channel
	Prometheus *setting.PrometheusSetting
)
