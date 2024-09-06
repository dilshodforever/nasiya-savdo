package handler

import (
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/kafka"
	r "github.com/dilshodforever/nasiya-savdo/storage/redis"
)

type Handler struct {
	User     pb.UserServiceClient
	redis    r.InMemoryStorageI
	producer kafka.KafkaProducer
}

func NewHandler(us pb.UserServiceClient, rdb r.InMemoryStorageI, pr kafka.KafkaProducer) *Handler {
	return &Handler{us, rdb, pr}
}
