package minio

import (
	"fmt"
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
	endpoint := "minio:9000"
	accessKeyID := "Dior"
	secretAccessKey := "isakov05@"
	useSSL := false

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connection to MinIO database established!")
	return &MinIO{
		client: minioClient,
	}, nil
}

func (m *MinIO) MinIO() storage.MinIOStorage {
	if m.minIO == nil {
		m.minIO = &MediaServiceServer{minioClient: m.client}
	}
	return m.minIO
}
