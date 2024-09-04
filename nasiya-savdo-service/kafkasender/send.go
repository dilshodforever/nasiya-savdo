package kafkasender

import (
	"encoding/json"
	"log"

	"github.com/dilshodforever/nasiya-savdo/kafka"
	"github.com/dilshodforever/nasiya-savdo/model"
)
func CreateNotification(kaf kafka.KafkaProducer, request model.Send) error {
	response, err := json.Marshal(request)
	if err != nil {
		log.Println("cannot produce messages via kafka", err.Error())
		return err
	}
	err = kaf.ProduceMessages("NasiaSendNatification", response)
	if err != nil {
		log.Fatal("Error while ProduceMessages: ", err.Error())
		return err
	}
	return nil
}
