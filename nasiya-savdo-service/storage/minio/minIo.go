package minio

import (
	"bytes"
	"context"
	"time"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MediaServiceServer struct {
	minioClient *minio.Client
}

func NewMediaServiceServer(minioClient *minio.Client) *MediaServiceServer {
	return &MediaServiceServer{
		minioClient: minioClient,
	}
}

func (s *MediaServiceServer) UploadFile(req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	file := bytes.NewReader(req.File)
	bucketName := "testbucket"
	location := "us-east-1"
	exists, err := s.minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to check if bucket exists")
	}

	if !exists {
		err = s.minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Failed to create bucket")
		}
	}

	objectName := "images/" + time.Now().Format("20060102150405") + "-" + req.Filename

	_, err = s.minioClient.PutObject(context.Background(), bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to upload image")
	}

	presignedURL, err := s.minioClient.PresignedGetObject(context.Background(), bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to generate URL")
	}

	return &pb.UploadFileResponse{Url: presignedURL.String()}, nil
}
