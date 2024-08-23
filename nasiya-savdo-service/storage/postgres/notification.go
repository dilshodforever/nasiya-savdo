package postgres

// import (
// 	"context"
// 	"log"

// 	chats "github.com/dilshodforever/fooddalivary-food/genprotos"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type ChatStorage struct {
// 	db *mongo.Database
// }

// func NewChatStorage(db *mongo.Database) *ChatStorage {
// 	return &ChatStorage{db: db}
// }

// // Chat yaratish funksiyasi
// func (c *ChatStorage) CreateChat(req *chats.CreateChatRequest) (*chats.CreateChatResponse, error) {

// 	coll := c.db.Collection("chats")
// 	_, err := coll.InsertOne(context.Background(), req)
// 	if err != nil {
// 		log.Printf("Failed to create chat: %v", err)
// 		return nil, err
// 	}

	
// 	return &chats.CreateChatResponse{
// 		Success: true,
// 	}, nil
// }

// // Xabar jo'natish funksiyasi
// func (c *ChatStorage) SendMessage(req *chats.SendMessageRequest) (*chats.SendMessageResponse, error) {

// 	coll := c.db.Collection("messages")
// 	_, err := coll.InsertOne(context.Background(), req)
// 	if err != nil {
// 		log.Printf("Failed to send message: %v", err)
// 		return nil, err
// 	}

// 	return &chats.SendMessageResponse{
// 		Success: true,
// 	}, nil
// }

// func (c *ChatStorage) GetMessages(req *chats.GetMessagesRequest) (*chats.GetMessagesResponse, error) {
// 	coll := c.db.Collection("messages")
// 	filter := bson.D{{Key: "chat_id", Value: req.ChatId}}

// 	cursor, err := coll.Find(context.Background(), filter)
// 	if err != nil {
// 		log.Printf("Failed to get messages: %v", err)
// 		return nil, err
// 	}
// 	defer cursor.Close(context.Background())

// 	var messages []*chats.Message
// 	for cursor.Next(context.Background()) {
// 		var message chats.Message
// 		if err := cursor.Decode(&message); err != nil {
// 			log.Printf("Failed to decode message: %v", err)
// 			return nil, err
// 		}
// 		messages = append(messages, &message)
// 	}

// 	if err := cursor.Err(); err != nil {
// 		log.Printf("Cursor error: %v", err)
// 		return nil, err
// 	}

// 	return &chats.GetMessagesResponse{
// 		Messages: messages,
// 	}, nil
// }
