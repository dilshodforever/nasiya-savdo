package kafkaconnect

import (
	"log"

	"github.com/dilshodforever/nasiya-savdo/kafka"
)

func ConnectToKafka() kafka.KafkaProducer{
	kaf, err := kafka.NewKafkaProducer([]string{"kafka:9092"})
	if err != nil {
		log.Fatal("Error while connection kafka: ", err.Error())
	}
	return kaf
}
