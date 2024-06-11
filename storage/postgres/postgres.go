package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"payment/config"
)

type Storage struct{
	Db *sql.DB
}

func DbConnection() (*Storage,error) {
	cfg := config.Load()
	con := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
	cfg.PostgresHost,cfg.PostgresUser,cfg.PostgresDatabase,cfg.PostgresPassword,cfg.PostgresPort)
	db,err := sql.Open("postgres",con)
	if err != nil{
		log.Fatal("Error while db connection",err)
		return nil,nil
	}
	err = db.Ping()
	if err != nil{
		log.Fatal("Error while db ping connection",err)
		return nil,nil
	}
	return &Storage{Db: db},nil
}