package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/5-oyimtixon/genprotos/notifications"
	"github.com/dilshodforever/5-oyimtixon/model"
	s "github.com/dilshodforever/5-oyimtixon/storage"
)

type NotificationService struct {
	stg s.InitRoot
	pb.UnimplementedNotificationtServiceServer
}

func NewNotificationService(stg s.InitRoot) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(req model.Send) error {
	err := s.stg.Notification().CreateNotification(req)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetAccountByid retrieves a notification by user_id
func (s *NotificationService) GetNotification(ctx context.Context, req *pb.GetNotificationByidRequest) (*pb.GetNotificationByidResponse, error) {
	notification, err := s.stg.Notification().GetNotification(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return notification, nil
}

// DeleteAccount deletes a notification by user_id
func (s *NotificationService) DeleteNotification(ctx context.Context, req *pb.GetNotificationByidRequest) (*pb.NotificationsResponse, error) {
	response,err := s.stg.Notification().DeleteNotification(req)
	if err != nil {
		log.Print(err)
		return response, err
	}

	return &pb.NotificationsResponse{
		Message: "Notification deleted successfully",
		Success: true,
	}, nil
}

// ListAccounts lists all notifications
func (s *NotificationService) ListNotification(ctx context.Context, req *pb.Void) (*pb.ListNotificationResponse, error) {
	notifications, err := s.stg.Notification().ListNotification(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}


	return notifications, nil
}
