package kafka

import (
	"encoding/json"
	"log"

	pb "github.com/dilshodforever/5-oyimtixon/model"
	"github.com/dilshodforever/5-oyimtixon/service"
)

func SendNotification(rootService *service.NotificationService) func(message []byte) {
	return func(message []byte) {
		var app pb.Send
		if err := json.Unmarshal(message, &app); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}
		err := rootService.CreateNotification(app)
		if err != nil {
			log.Printf("Cannot create evaluation via Kafka: %v", err)
			return
		}
		log.Print("Created evaluation")
	}
}