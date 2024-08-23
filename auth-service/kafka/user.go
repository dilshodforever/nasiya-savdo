package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "gitlab.com/lingualeap/auth/genprotos/users"
	"gitlab.com/lingualeap/auth/service"
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
