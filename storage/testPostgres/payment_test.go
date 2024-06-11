package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	pb "github.com/Project_Restaurant/Payment-Service/genproto/payment"
	"github.com/Project_Restaurant/Payment-Service/storage/postgres"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "sayyidmuhammad"
	dbPassword = "root"
	dbName     = "payment"
)

func newTestDB(t *testing.T) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}

	return db
}

func TestPaymentRepo(t *testing.T) {
	db := newTestDB(t)
	defer db.Close()
	repo := postgres.NewPaymentRepo(db)

	t.Run("CreatePayment", func(t *testing.T) {
		req := &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        100.00,
				PaymentMethod: "card",
				PaymentStatus: "success",
			},
		}
		payment, err := repo.CreatePayment(context.Background(), req)
		if err != nil {
			t.Fatal(err)
		}
		assert.NotEmpty(t, payment.Id)
		assert.Equal(t, req.Payment.ReservationId, payment.ReservationId)
		assert.Equal(t, req.Payment.Amount, payment.Amount)
		assert.Equal(t, req.Payment.PaymentMethod, payment.PaymentMethod)
		assert.Equal(t, req.Payment.PaymentStatus, payment.PaymentStatus)
		// Add assertions for CreatedAt and UpdatedAt
	})

	t.Run("GetPaymentById", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        150.50,
				PaymentMethod: "paypal",
				PaymentStatus: "pending",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		payment, err := repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, createdPayment.Id, payment.Id)
		assert.Equal(t, createdPayment.ReservationId, payment.ReservationId)
		assert.Equal(t, createdPayment.Amount, payment.Amount)
		assert.Equal(t, createdPayment.PaymentMethod, payment.PaymentMethod)
		assert.Equal(t, createdPayment.PaymentStatus, payment.PaymentStatus)
		// Add assertions for CreatedAt and UpdatedAt
	})

	t.Run("UpdatePayment", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        200.00,
				PaymentMethod: "cash",
				PaymentStatus: "failed",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		updateReq := &pb.UpdatePaymentRequest{
			Payment: &pb.Payment{
				Id:            createdPayment.Id,
				ReservationId: uuid.NewString(),
				Amount:        250.00,
				PaymentMethod: "credit card",
				PaymentStatus: "success",
			},
		}

		updatedPayment, err := repo.UpdatePayment(context.Background(), updateReq)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, updateReq.Payment.Id, updatedPayment.Id)
		assert.Equal(t, updateReq.Payment.ReservationId, updatedPayment.ReservationId)
		assert.Equal(t, updateReq.Payment.Amount, updatedPayment.Amount)
		assert.Equal(t, updateReq.Payment.PaymentMethod, updatedPayment.PaymentMethod)
		assert.Equal(t, updateReq.Payment.PaymentStatus, updatedPayment.PaymentStatus)
		// Add assertions for UpdatedAt
	})

	t.Run("DeletePayment", func(t *testing.T) {
		createdPayment, err := repo.CreatePayment(context.Background(), &pb.CreatePaymentRequest{
			Payment: &pb.Payment{
				ReservationId: uuid.NewString(),
				Amount:        300.00,
				PaymentMethod: "bank transfer",
				PaymentStatus: "refunded",
			},
		})
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.DeletePayment(context.Background(), &pb.DeletePaymentRequest{Id: createdPayment.Id})
		if err != nil {
			t.Fatal(err)
		}

		_, err = repo.GetPaymentById(context.Background(), &pb.GetPaymentRequest{Id: createdPayment.Id})
		assert.NotNil(t, err) // Expect an error since the payment should be deleted
	})
}
