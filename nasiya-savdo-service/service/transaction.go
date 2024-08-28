package service

import (
	"context"
	"log"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	s "github.com/dilshodforever/nasiya-savdo/storage"
)

type TransactionService struct {
	stg s.InitRoot
	pb.UnimplementedTransactionServiceServer
}

func NewTransactionService(stg s.InitRoot) *TransactionService {
	return &TransactionService{stg: stg}
}

func (s *TransactionService) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	resp, err := s.stg.Transaction().CreateTransaction(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *TransactionService) GetTransaction(ctx context.Context, req *pb.TransactionIdRequest) (*pb.GetTransactionResponse, error) {
	resp, err := s.stg.Transaction().GetTransaction(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *TransactionService) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error) {
	resp, err := s.stg.Transaction().UpdateTransaction(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}
func (s *TransactionService) DeleteTransaction(ctx context.Context, req *pb.TransactionIdRequest) (*pb.TransactionResponse, error) {
	resp, err := s.stg.Transaction().DeleteTransaction(req)
	if err != nil {
		log.Print(err)
		return resp, err
	}
	return resp, nil
}

func (s *TransactionService) ListTransactions(ctx context.Context, req *pb.GetAllTransactionRequest) (*pb.GetAllTransactionResponse, error) {
	resp, err := s.stg.Transaction().ListTransactions(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}

func (s *TransactionService) CheckTransactions(ctx context.Context,req *pb.CheckRequest) (*pb.CheckResponse, error) {
	resp, err := s.stg.Transaction().CheckTransactions(req)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return resp, nil
}


