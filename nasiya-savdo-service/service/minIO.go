package service

import (
	"context"

	"github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/storage"
)

type MiiIOService struct {
	genprotos.UnimplementedMediaServiceServer
	storage.MinIoRoot
}

func NewMinIOService(stg storage.MinIoRoot) *MiiIOService {
	return &MiiIOService{
		MinIoRoot: stg,
	}
}

func (s *MiiIOService) UploadFile(contex context.Context, req *genprotos.UploadFileRequest) (*genprotos.UploadFileResponse, error) {
	resp, err := s.MinIO().UploadFile(contex, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
