package service

import (
	"context"
	"log"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"
	"github.com/Project_Restaurant/Payment-Service/storage/postgres"
)

type PaymentService struct {
	stg postgres.Storage
	pb.UnimplementedPaymentServiceServer
}

func NewPAymentService(stg postgres.Storage) *PaymentService {
	return &PaymentService{stg: stg}
}

func (ps *PaymentService) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.Payment, error) {
	payment, err := ps.stg.Payments.CreatePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while create service", err)
		return &pb.Payment{}, err
	}
	return payment, nil
}

func (ps *PaymentService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.Payment, error) {
	payment, err := ps.stg.Payments.UpdatePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while update service", err)
		return &pb.Payment{}, err
	}
	return payment, nil
}

func (ps *PaymentService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.Payment, error) {
	payment, err := ps.stg.Payments.GetPaymentById(ctx, req)
	if err != nil {
		log.Fatal("Error while update service", err)
		return &pb.Payment{}, err
	}
	return payment, nil
}

func (ps *PaymentService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	payment, err := ps.stg.Payments.DeletePayment(ctx, req)
	if err != nil {
		log.Fatal("Error while update service", err)
		return &pb.DeletePaymentResponse{Message: "Error"}, err
	}
	return payment, nil
}
