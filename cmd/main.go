package main

import (
	"log"
	"net"
	pb "payment/genproto/payment"
	"payment/service"
	"payment/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.DbConnection()
	if err != nil {
		log.Fatal("Error while db connection", err)
		return
	}

	newServer := grpc.NewServer()

	pb.RegisterPaymentServiceServer(newServer, service.NewPAymentService(*db))

	lis, err := net.Listen("tcp", "50050")
	if err != nil {
		log.Fatal("Error while listen tcp", err)
		return
	}
	
	err = newServer.Serve(lis)
	if err != nil {
		log.Fatal("Error while newServe", err)
	}
}
