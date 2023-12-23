package main

import (
	"context"
	"fmt"
	"log"

	healthv1 "github.com/mi11km/neuron-visualizer/server/proto/health/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	address := "localhost:8080"
	conn, err := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	client := healthv1.NewHealthServiceClient(conn)
	res, err := client.Check(context.Background(), &healthv1.CheckRequest{
		Service: "NeuronService",
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", res)
}
