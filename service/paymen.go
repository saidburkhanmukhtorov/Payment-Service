package service

import (
	"context"
	"log"

	"github.com/Diyor/Project/Payment/genproto/payment"
	"github.com/Diyor/Project/Payment/storage/postgres"
)

type PaymenService struct {
	p postgres.Storage
	payment.UnimplementedPaymentServiceServer
}

func NewPaymenService(p postgres.Storage) *PaymenService {
	return &PaymenService{p: p}
}

func (s *PaymenService) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	_, err := s.p.Payments.CreatePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while creating payment")
		return &payment.CreatePaymentResponse{}, err
	}
	return &payment.CreatePaymentResponse{}, nil
}

func (s *PaymenService) GetPayment(ctx context.Context, req *payment.GetPaymentRequest) (*payment.GetPaymentResponse, error){
	payment, err := s.p.Payments.GetPaymentById(ctx, req)
	if err != nil {
		log.Fatal("Error while getting payment")
		return nil, err
	}
	return payment, nil
}

func (s *PaymenService) UpdatePayment(ctx context.Context, req *payment.UpdatePaymentRequest) (*payment.UpdatePaymentResponse, error){
	payment, err := s.p.Payments.UpdatePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while updating payment")
		return nil, err
	}
	return payment, nil
}


func (s *PaymenService) DeletePayment(ctx context.Context, req *payment.DeletePaymentRequest) (*payment.DeletePaymentResponse, error){	
	payment, err := s.p.Payments.DeletePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while deleting payment")
		return nil, err
	}
	return payment, nil
}