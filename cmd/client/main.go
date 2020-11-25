package main

import (
	"context"
	"log"

	calcul "github.com/idirall22/grpc_calcul/api/pb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not dial server: %v", err)
	}
	defer cc.Close()

	client := calcul.NewCalculServiceClient(cc)
	res, err := client.Add(context.Background(), &calcul.AddReq{A: 5, B: 4})
	if err != nil {
		log.Fatalf("Could not dial server: %v", err)
	}
	log.Println("Response:", res.Result)
}
