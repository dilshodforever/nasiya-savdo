package minio

import (
	"log"

	"github.com/dilshodforever/nasiya-savdo/storage"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinIO struct {
	client *minio.Client
	minIO  storage.MinIOStorage
}

func InitMinioClient() (storage.MinIoRoot, error) {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "pen0FH8fIYhbCcT0zCGK"
	secretAccessKey := "q73ArcpEWn3bypibpFu0wcW7rsC6kyHQ8u9yOVUD"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &MinIO{
		client: minioClient,
	}, nil
}


func (m *MinIO) MinIO() storage.MinIOStorage {
	if m.minIO != nil {
		m.minIO = &MediaServiceServer{minioClient: m.client}
	}
	return m.minIO
}
