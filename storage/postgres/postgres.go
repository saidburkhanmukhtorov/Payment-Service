package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Diyor/Project/Payment/storage"
	_ "github.com/lib/pq" 
)

type Storage struct {
	Db       *sql.DB
	Payments storage.PaymentI
}

func ConnectDB() (*Storage, error) {
	psql := "user=postgres password=20005 dbname=restarount sslmode=disable"
	db, err := sql.Open("postgres", psql)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	payment := NewPaymentRepo(db)
	return &Storage{
		Db:       db,
		Payments: payment,
	}, nil
}

func (c *Storage) Payment() storage.PaymentI {
	if c.Payments == nil {
		c.Payments = NewPaymentRepo(c.Db)
	}
	return c.Payments
}
