package minio

import (
	"bytes"
	"context"
	"fmt"
	"log"
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

func (s *MediaServiceServer) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	fmt.Println("Salom")
	file := bytes.NewReader(req.File)
	bucketName := "testbucket"
	location := "us-east-1"
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to check if bucket exists")
	}

	if !exists {
		err = s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
		if err != nil {
			log.Println("Failed to create bucket")
			return nil, status.Errorf(codes.Internal, "Failed to create bucket")
		}
	}

	objectName := "images/" + time.Now().Format("20060102150405") + "-" + req.Filename

	_, err = s.minioClient.PutObject(ctx, bucketName, objectName, file, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Println("Failed to upload image")
		return nil, status.Errorf(codes.Internal, "Failed to upload image")
	}

	presignedURL, err := s.minioClient.PresignedGetObject(ctx, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		log.Println("Failed to generate URL")
		return nil, status.Errorf(codes.Internal, "Failed to generate URL")
	}

	return &pb.UploadFileResponse{Url: presignedURL.String()}, nil
}
