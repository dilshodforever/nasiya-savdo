package service

import (
	"context"

	"github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/storage"
)

type miiIOService struct {
	genprotos.UnimplementedMediaServiceServer
	storage.MinIoRoot
}

func NewMinIOService(stg storage.MinIoRoot) *miiIOService {
	return &miiIOService{
		MinIoRoot: stg,
	}
}

func (s *miiIOService) UploadFile(contex context.Context, req *genprotos.UploadFileRequest) (*genprotos.UploadFileResponse, error) {
	resp, err := s.MinIO().UploadFile(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
