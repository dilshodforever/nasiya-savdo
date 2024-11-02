package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	s "github.com/dilshodforever/nasiya-savdo/storage"
)

type ExchangeService struct {
	stg s.InitRoot
	pb.UnimplementedExchangeServiceServer
}

func NewExchangeService(stg s.InitRoot) *ExchangeService {
	return &ExchangeService{stg: stg}
}

func (s *ExchangeService) CreateExchange(ctx context.Context, req *pb.CreateExchangeRequest) (*pb.ExchangeResponse, error) {
	resp, err := s.stg.Exchange().CreateExchange(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ExchangeService) GetExchange(ctx context.Context, req *pb.ExchangeIdRequest) (*pb.GetExchangeResponse, error) {
	resp, err := s.stg.Exchange().GetExchange(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *ExchangeService) UpdateExchange(ctx context.Context, req *pb.UpdateExchangeRequest) (*pb.ExchangeResponse, error) {
	resp, err := s.stg.Exchange().UpdateExchange(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ExchangeService) DeleteExchange(ctx context.Context, req *pb.ExchangeIdRequest) (*pb.ExchangeResponse, error) {
	resp, err := s.stg.Exchange().DeleteExchange(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *ExchangeService) ListExchanges(ctx context.Context, req *pb.GetAllExchangeRequest) (*pb.GetAllExchangeResponse, error) {
	resp, err := s.stg.Exchange().ListExchanges(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *ExchangeService) GetStatistika(ctx context.Context,req *pb.ExchangeStatisticsRequest) (*pb.ExchangeStatisticsResponse, error) {
	resp, err := s.stg.Exchange().GetMonthlyStatistics(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *ExchangeService) GetExchangeGetbyProductId(ctx context.Context,req *pb.GetExchangeGetbyProductIdRequest) (*pb.GetExchangeGetbyProductIdResponse, error) {
	resp, err := s.stg.Exchange().GetExchangeGetbyProductId(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}


