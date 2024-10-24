package postgres

import (
	ctx "context"
	"fmt"
	"log"
	"time"

	pb "github.com/dilshodforever/5-oyimtixon/genprotos"
	"github.com/dilshodforever/5-oyimtixon/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountService struct {
	db *mongo.Database
}

func NewAccountService(db *mongo.Database) *AccountService {
	return &AccountService{db: db}
}

func (s *AccountService) CreateNotification(req model.Send) error {
	coll := s.db.Collection("notifications")
	id := uuid.NewString()
	_, err := coll.InsertOne(ctx.Background(), bson.M{
		"id":      id,
		"UserId":  req.Userid,
		"message": req.Message,
		"ContractId":req.ContractId,
		"CreatedAt":time.Now().String(),
		"IsRead": false,
	})
	if err != nil {
		log.Printf("Failed to create account: %v", err)
		return err
	}
	return nil
}

func (s *AccountService) GetNotification(req *pb.GetNotificationByidRequest) (*pb.GetNotificationByidResponse, error) {
	coll := s.db.Collection("notifications")
	var notification model.Send
	err := coll.FindOne(ctx.Background(), bson.M{"UserId": req.UserId}).Decode(&notification)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("sizda hali notificationlar mavjud emas")
		}
		log.Printf("Failed to retrieve notification: %v", err)
		return nil, err
	}
	err=s.MarkNotificationAsRead(req.UserId)
	if err !=nil {
		return nil, fmt.Errorf("error while update notification's isread")
	}
	return &pb.GetNotificationByidResponse{
		UserId:  notification.Userid,
		Message: notification.Message,
		ContractId: notification.ContractId,
		CreatedAt: notification.CreatedAt,
	}, nil
}

// DeleteAccount deletes a notification by user_id
func (s *AccountService) DeleteNotification(req *pb.GetNotificationByidRequest) (*pb.NotificationsResponse, error) {
	coll := s.db.Collection("notifications")
	res, err := coll.DeleteOne(ctx.Background(), bson.M{"UserId": req.UserId})
	if err != nil {
		log.Printf("Failed to delete notification: %v", err)
		return &pb.NotificationsResponse{
			Message: "Error while deleting notification",
			Success: false,
		}, err
	}

	if res.DeletedCount == 0 {
		return &pb.NotificationsResponse{
			Message: "No notification found to delete",
			Success: false,
		}, nil
	}

	return &pb.NotificationsResponse{
		Message: "Notification deleted successfully",
		Success: true,
	}, nil
}


func (s *AccountService) ListNotification(req *pb.GetNotificationByidRequest) (*pb.ListNotificationResponse, error) {
    coll := s.db.Collection("notifications")
    
    // Create a filter to match the requested user ID
    filter := bson.M{"UserId": req.UserId}
    
    // Find documents that match the filter
    cursor, err := coll.Find(ctx.Background(), filter)
    if err != nil {
        log.Printf("Failed to list notifications: %v", err)
        return nil, err
    }
    defer cursor.Close(ctx.Background())

    var notifications []*pb.GetNotificationByidResponse
    for cursor.Next(ctx.Background()) {
        var notification model.Send
        if err := cursor.Decode(&notification); err != nil {
            log.Printf("Failed to decode notification: %v", err)
            return nil, err
        }
        notifications = append(notifications, &pb.GetNotificationByidResponse{
            UserId:     notification.Userid,
            Message:    notification.Message,
            ContractId: notification.ContractId,
            CreatedAt:  notification.CreatedAt,
        })
    }

    if err := cursor.Err(); err != nil {
        log.Printf("Cursor error: %v", err)
        return nil, err
    }

    return &pb.ListNotificationResponse{Notifications: notifications}, nil
}


func (s *AccountService) MarkNotificationAsRead(userid string) error {
    coll := s.db.Collection("notifications")

    // Create a filter to match the notification by UserId and NotificationId
    filter := bson.M{"UserId": userid}

    // Define the update: set IsRead = true
    update := bson.M{
        "$set": bson.M{
            "IsRead": true,
        },
    }

    // Perform the update
    _, err := coll.UpdateOne(ctx.Background(), filter, update)
    if err != nil {
        log.Printf("Failed to mark notification as read: %v", err)
        return err
    }

    return nil
}
