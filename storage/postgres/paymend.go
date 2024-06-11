package postgres

import (
	"context"
	"database/sql"
	"log"
	pb "payment/genproto/payment"

	"github.com/google/uuid"
)

type PaymentRepo struct {
	Db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{Db: db}
}

func (db *PaymentRepo) CreatePayment(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.Payment, error) {
	id := uuid.New().String()
	query := `
		INSERT INTO 
			restaurants (
				id,
				reservation_id,
				amount,
				payment_method,
				payment_status,
			) 
		VALUES (
				$1, 
				$2, 
				$3,
				$4,
				$5
			)
		RETURNING 
			id, 
			reservation_id,
			amount,
			payment_method,
			payment_status,
			created_at,
			updated_at
	`
	pay := pb.Payment{}
	err := db.Db.QueryRow(query, id, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus).Scan(
		&pay.Id,
		&pay.ReservationId,
		&pay.Amount,
		&pay.PaymentMethod,
		&pay.PaymentStatus,
		&pay.CreatedAt,
		&pay.UpdatedAt,
	)
	if err != nil {
		log.Fatal("Error while create payment in postgres", err)
		return nil, nil
	}

	return &pay, nil
}

func (db *PaymentRepo) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.Payment, error) {
	query := `
		update payment set reservation_id = $1,
		amount = $2,
		payment_method = $3,
		payment_status = $4,
		updated_at = NOW()
		where id = $5
	`
	_, err := db.Db.Exec(query, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus, req.Payment.Id)
	if err != nil {
		log.Fatal("error while update Payment in postgres", err)
		return nil, nil
	}
	pay, err := db.GetPaymentById(ctx, &pb.GetPaymentRequest{Id: req.Payment.Id})
	if err != nil {
		log.Fatal("Error while update payment error get payment", err)
		return nil, nil
	}
	return pay, nil
}

func (db *PaymentRepo) GetPaymentById(ctx context.Context, req *pb.GetPaymentRequest) (*pb.Payment, error) {
	query := `
		select 
			id, 
			reservation_id,
			amount,
			payment_method,
			payment_status,
			created_at,
			updated_at
		from payment
		where id = $1 and deleted_at = 0
	`
	pay := pb.Payment{}
	err := db.Db.QueryRow(query, req.Id).Scan(
		&pay.Id,
		&pay.Amount,
		&pay.PaymentMethod,
		&pay.PaymentStatus,
		&pay.CreatedAt,
		&pay.UpdatedAt,
	)
	if err != nil {
		log.Fatal("Error while get payment in postgres", err)
		return nil, nil
	}
	return &pay, nil
}

func (db *PaymentRepo) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.DeletePaymentResponse, error) {
	query := `
		update payment set deleted_at = extract(epoch from NOW()) where id = $1 and deleted_at = 0
	`
	_, err := db.Db.Exec(query, req.Id)
	if err != nil {
		log.Fatal("Error while deleted payment", err)
		return nil, nil
	}
	return &pb.DeletePaymentResponse{Message: "Payment deleted"}, nil
}
