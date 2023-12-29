package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	"github.com/mi11km/neuron-visualizer/server/interfaces"
	healthv1 "github.com/mi11km/neuron-visualizer/server/proto/health/v1"
	neuronv1 "github.com/mi11km/neuron-visualizer/server/proto/neuron/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 8080

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	neuronServiceServer, err := interfaces.NewNeuronServiceServer("simulations")
	if err != nil {
		panic(err)
	}
	healthv1.RegisterHealthServiceServer(s, &interfaces.HealthServiceServer{})
	neuronv1.RegisterNeuronServiceServer(s, neuronServiceServer)

	reflection.Register(s)

	go func() {
		slog.Info(fmt.Sprintf("start gRPC server port: %v", port))
		if err := s.Serve(listener); err != nil {
			slog.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	slog.Info("stopping gRPC server...")
	s.GracefulStop()
}
