package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	s "github.com/dilshodforever/nasiya-savdo/storage"
)

type ContractService struct {
	stg s.InitRoot
	pb.UnimplementedContractServiceServer
}

func NewContractService(stg s.InitRoot) *ContractService {
	return &ContractService{stg: stg}
}

func (s *ContractService) CreateContract(ctx context.Context, req *pb.CreateContractRequest) (*pb.ContractResponse, error) {
	resp, err := s.stg.Contract().CreateContract(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ContractService) GetContract(ctx context.Context, req *pb.ContractIdRequest) (*pb.GetContractResponse, error) {
	resp, err := s.stg.Contract().GetContract(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *ContractService) UpdateContract(ctx context.Context, req *pb.UpdateContractRequest) (*pb.ContractResponse, error) {
	resp, err := s.stg.Contract().UpdateContract(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ContractService) DeleteContract(ctx context.Context, req *pb.ContractIdRequest) (*pb.ContractResponse, error) {
	resp, err := s.stg.Contract().DeleteContract(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ContractService) ListContracts(ctx context.Context, req *pb.GetAllContractRequest) (*pb.GetAllContractResponse, error) {
	resp, err := s.stg.Contract().ListContracts(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}
