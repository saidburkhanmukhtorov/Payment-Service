package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/Diyor/Project/Payment/genproto/payment"
	"github.com/google/uuid"
)

type PaymentRepo struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepo {
	return &PaymentRepo{db: db}
}

func (p *PaymentRepo) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	res := payment.Payment{}
	var createdAt, updatedAt time.Time
	id := uuid.New().String()
	query := `
	insert into payments 
	(id, reservation_id, amout, payment_method, payment_status)
	values ($1, $2, $3, $4, $5)`
	err := p.db.QueryRowContext(ctx, query, id, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentMethod, req.Payment.PaymentStatus).
		Scan(res.Id, res.ReservationId, res.Amount, res.PaymentMethod, res.PaymentStatus)
	if err != nil {
		slog.Error("failed to create payment")
		return nil, err
	}
	res.CreatedAt = createdAt.Format(time.RFC3339)
	res.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &payment.CreatePaymentResponse{Payment: &res}, nil
}


func (p *PaymentRepo) UpdatePayment(ctx context.Context, req *payment.UpdatePaymentRequest) (*payment.UpdatePaymentResponse, error) {
	query := `
	update payments set
		reservation_id = $1,
		amount = $2,
		payment_status = $3,
		payment_method = $4,
		updated_at = now()
	where id = $5`
	
	_, err := p.db.ExecContext(ctx, query, req.Payment.ReservationId, req.Payment.Amount, req.Payment.PaymentStatus, req.Payment.PaymentMethod, req.Payment.Id)
	if err != nil {
		slog.Error("failed to update payment", "error", err)
		return nil, err
	}

	return &payment.UpdatePaymentResponse{Payment: req.Payment}, nil
}



func (p *PaymentRepo) GetPaymentById(ctx context.Context, req *payment.GetPaymentRequest) (*payment.GetPaymentResponse, error){
	res := payment.Payment{}
	var createdAt, updatedAt time.Time
	query := `
	select * from payments where id = $1`

	err := p.db.QueryRowContext(ctx, query, req.Id).
	Scan(&res.Id, &res.ReservationId, &res.Amount, &res.PaymentMethod, &res.PaymentStatus, &res.CreatedAt, &res.UpdatedAt)
	if err != nil{
		slog.Error("failed to get payment")
		return nil, err
	}
	res.CreatedAt = createdAt.Format(time.RFC3339)
	res.UpdatedAt = updatedAt.Format(time.RFC3339)
	return &payment.GetPaymentResponse{Payment: &res}, nil
}


func (p *PaymentRepo) DeletePayment(ctx context.Context, req *payment.DeletePaymentRequest) (*payment.DeletePaymentResponse, error){
	query := `delete from payments where id = $1`
	_, err := p.db.ExecContext(ctx, query, req.Id)
	if err != nil{
		slog.Error("failed to delete payment")
		return nil, err
	}
	return &payment.DeletePaymentResponse{Message: "payment deleted"}, nil
}
