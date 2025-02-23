package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	kafka "github.com/segmentio/kafka-go"
	"strings"
	"time"
)

var (
	kafkaProducer *kafka.Writer
)

const (
	kafkaURL   = "localhost:19094"
	kafkaTopic = "user_topic_vip"
	groupID    = "test-group"
)

// for producer
func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{}, // can bang tai mac dinh
	}
}

// for consumer
func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        brokers, // list of broker addresses []string{"localhost:1", "localhost:2", ...}
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       10e3,            // 10KB   // 10KB - the minimum number of bytes to fetch in a request
		MaxBytes:       10e6,            // 10MB    // 10MB - the maximum number of bytes to fetch in a request
		CommitInterval: 1 * time.Second, // the interval between which the reader commits the offset of messages it has read
		StartOffset:    kafka.FirstOffset,
	})
}

type StockInfo struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

// mua ban chung khoang
// new stock
func newStock(msg, typeMsg string) *StockInfo {
	s := StockInfo{}
	s.Message = msg
	s.Type = typeMsg

	return &s
}

// action
func actionStock(c *gin.Context) {
	s := newStock(c.Query("msg"), c.Query("type"))
	body := make(map[string]interface{})
	body["action"] = "action"
	body["info"] = s

	jsonBody, _ := json.Marshal(body)
	// tao  message cua kafka
	msg := kafka.Message{
		Key:   []byte("action"),
		Value: []byte(jsonBody),
	}
	// viet message by producer
	err := kafkaProducer.WriteMessages(context.Background(), msg)
	if err != nil {
		c.JSON(200, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"err": "",
		"msg": "action successfully",
	})

}

// consumer hóng mua ATC
func RegisterConsumerATC(id int) {
	// group consumer??
	// kafkaGroupId := "consumer-group-"
	kafkaGroupId := fmt.Sprintf("consumer-group-%d", id) // "consumer-group-1"
	reader := getKafkaReader(kafkaURL, kafkaTopic, kafkaGroupId)

	defer reader.Close()

	fmt.Printf("Consumer(%d) Hong Phien ATC::\n", id)
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Printf("Consumer(%d) error: %v, id, err")

		}
		fmt.Printf("Consumer(%d), hong topic:%v, partition:%v, offset:%v, time: %d %s = %s\n", id, m.Topic, m.Partition, m.Offset, m.Time.Unix(), string(m.Key), string(m.Value))

	}

}

func main() {
	r := gin.Default()
	kafkaProducer = getKafkaWriter(kafkaURL, kafkaTopic)
	defer kafkaProducer.Close()

	r.POST("action/stock", actionStock)

	// đăng ký 2 user để mua stock trong ATC (1) (2)
	go RegisterConsumerATC(1)
	go RegisterConsumerATC(2)
	//go RegisterConsumerATC(3)
	//go RegisterConsumerATC(4)

	r.Run(":8999")

}
