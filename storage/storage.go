package storage

import (
	"context"

	pb "github.com/Diyor/Project/Payment/genproto/payment"
)


type StorageI interface {
	Payment() PaymentI
}

type PaymentI interface {
	CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error)
	UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.UpdatePaymentResponse, error)
	GetPaymentById(ctx context.Context, req *pb.GetPaymentRequest) (*pb.GetPaymentResponse, error)
	DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error)
}
