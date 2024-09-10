package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/service"
)

func UserCreateHandler(userServ *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.UserReq
		if err := json.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userServ.Register(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create User via Kafka: %v", err)
			return
		}
		log.Printf("Created User: %+v", res)
	}
}

func UserUpdateHandler(userServ *service.UserService) func(message []byte) {
	return func(message []byte) {
		var user pb.User
		if err := json.Unmarshal(message, &user); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		res, err := userServ.Update(context.Background(), &user)
		if err != nil {
			log.Printf("Cannot create User via Kafka: %v", err)
			return
		}
		log.Printf("Created User: %+v", res)
	}
}
