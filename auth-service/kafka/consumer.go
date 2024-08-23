package kafka

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumerManager struct {
	consumers map[string]*kafka.Reader
	handlers  map[string]func(message []byte)
	mu        sync.Mutex
}

func NewKafkaConsumerManager() *KafkaConsumerManager {
	return &KafkaConsumerManager{
		consumers: make(map[string]*kafka.Reader),
		handlers:  make(map[string]func(message []byte)),
	}
}

var ErrConsumerAlreadyExists = errors.New("consumer for this topic already exists")

func (kcm *KafkaConsumerManager) RegisterConsumer(brokers []string, topic, groupID string, handler func(message []byte)) error {
	kcm.mu.Lock()
	defer kcm.mu.Unlock()

	if _, exists := kcm.consumers[topic]; exists {
		return ErrConsumerAlreadyExists
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})
	kcm.consumers[topic] = reader
	kcm.handlers[topic] = handler

	go kcm.consumeMessages(topic)

	return nil
}

func (kcm *KafkaConsumerManager) consumeMessages(topic string) {
	reader := kcm.consumers[topic]
	handler := kcm.handlers[topic]

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message from topic %s: %v", topic, err)
			continue
		}
		handler(msg.Value)
	}
}

func (kcm *KafkaConsumerManager) Close() error {
	kcm.mu.Lock()
	defer kcm.mu.Unlock()

	for _, reader := range kcm.consumers {
		if err := reader.Close(); err != nil {
			return err
		}
	}
	return nil
}
