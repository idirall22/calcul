package main

import (
	"log"
	"net"

	calcul "github.com/idirall22/grpc_calcul/api/pb"
	calculservice "github.com/idirall22/grpc_calcul/pkg"
	"google.golang.org/grpc"
)

var add = ":8081"

func main() {
	l, err := net.Listen("tcp", add)
	if err != nil {
		log.Fatalf("Could not crea	te a listner: %v", err)
	}
	defer l.Close()

	srv := grpc.NewServer()
	calcul.RegisterCalculServiceServer(srv, calculservice.NewService())

	log.Println("Server runing on", add)
	srv.Serve(l)
}
