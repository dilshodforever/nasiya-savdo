package storage

import ("github.com/dilshodforever/5-oyimtixon/model"
pb "github.com/dilshodforever/5-oyimtixon/genprotos/notifications"
)

type InitRoot interface {
	Notification() NotificationService
}

type NotificationService interface {
	CreateNotification(req model.Send) error
	GetNotification(req *pb.GetNotificationByidRequest) (*pb.GetNotificationByidResponse, error)
	DeleteNotification(req *pb.GetNotificationByidRequest) (*pb.NotificationsResponse, error) 
	ListNotification(req *pb.Void) (*pb.ListNotificationResponse, error)
}
