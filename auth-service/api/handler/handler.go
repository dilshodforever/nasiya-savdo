package handler

import (
	pb "gitlab.com/lingualeap/auth/genprotos/users"
	"gitlab.com/lingualeap/auth/kafka"
	r "gitlab.com/lingualeap/auth/storage/redis"
)

type Handler struct {
	User     pb.UserServiceClient
	redis    r.InMemoryStorageI
	producer kafka.KafkaProducer
}

func NewHandler(us pb.UserServiceClient, rdb r.InMemoryStorageI, pr kafka.KafkaProducer) *Handler {
	return &Handler{us, rdb, pr}
}
