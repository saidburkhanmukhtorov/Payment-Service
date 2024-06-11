package storage

import (
	"context"
	pb "payment/genproto/payment"
)

type StorageI interface {
	Payment() PaymentI
}

type PaymentI interface {
	CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.Payment, error)
	UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.Payment, error)
	GetPaymentById(ctx context.Context, req *pb.GetPaymentRequest) (*pb.Payment, error)
	DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error)
}
