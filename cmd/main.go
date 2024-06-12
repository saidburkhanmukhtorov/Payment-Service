package main

import (
	"log"
	"net"

	"github.com/Diyor/Project/Payment/genproto/payment"
	"github.com/Diyor/Project/Payment/service"
	"github.com/Diyor/Project/Payment/storage/postgres"
	"google.golang.org/grpc"
)



func main(){
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("failed to connect database")
		return
	}
	server := grpc.NewServer()


	payment.RegisterPaymentServiceServer(server, service.NewPaymenService(*db))
	lis, err := net.Listen( "tcp", ":8080")
	if err != nil {
		log.Fatal("failed to listen")
		return
	}
	log.Printf("server listening at %v", lis.Addr())
	log.Println("server starting...")
	server.Serve(lis)

}