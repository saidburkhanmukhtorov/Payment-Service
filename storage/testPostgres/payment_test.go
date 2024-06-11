package postgres

import (
	"context"
	"testing"

	pb "payment/genproto/payment"
	"payment/storage/postgres"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func NewTestPaymentRepo(t *testing.T) *postgres.PaymentRepo {
	db, err := postgres.DbConnection()
	if err != nil {
		t.Fatal("err while connecting db", err)
	}
	return postgres.NewPaymentRepo(db.Db)
}

func TestCreatePayment(t *testing.T) {
	repo := NewTestPaymentRepo(t)
	req := &pb.CreatePaymentRequest{
		Payment: &pb.Payment{
			ReservationId: uuid.NewString(),
			Amount:        100,
			PaymentMethod: "credit_card",
			PaymentStatus: "pending",
		},
	}
	resp, err := repo.CreatePayment(context.Background(), req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Id)
	assert.Equal(t, resp.Id, req.Payment.ReservationId)
	assert.Equal(t, req.Payment.Amount, resp.Amount)
	assert.Equal(t, req.Payment.PaymentMethod, resp.PaymentMethod)
	assert.Equal(t, req.Payment.PaymentStatus, resp.PaymentStatus)
}
