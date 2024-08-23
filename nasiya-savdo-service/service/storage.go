package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	s "github.com/dilshodforever/nasiya-savdo/storage"
)

type StorageService struct {
	stg s.InitRoot
	pb.UnimplementedStorageServiceServer
}

func NewStorageService(stg s.InitRoot) *StorageService {
	return &StorageService{stg: stg}
}

func (s *StorageService) CreateStorage(ctx context.Context, req *pb.CreateStorageRequest) (*pb.StorageResponse, error) {
	resp, err := s.stg.Storage().CreateStorage(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *StorageService) GetStorage(ctx context.Context, req *pb.StorageIdRequest) (*pb.GetStorageResponse, error) {
	resp, err := s.stg.Storage().GetStorage(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *StorageService) UpdateStorage(ctx context.Context, req *pb.UpdateStorageRequest) (*pb.StorageResponse, error) {
	resp, err := s.stg.Storage().UpdateStorage(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *StorageService) DeleteStorage(ctx context.Context, req *pb.StorageIdRequest) (*pb.StorageResponse, error) {
	resp, err := s.stg.Storage().DeleteStorage(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *StorageService) ListStorages(ctx context.Context, req *pb.GetAllStorageRequest) (*pb.GetAllStorageResponse, error) {
	resp, err := s.stg.Storage().ListStorages(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
