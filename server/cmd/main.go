package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"

	healthv1 "github.com/mi11km/neuron-visualizer/server/proto/health/v1"
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

	healthv1.RegisterHealthServiceServer(s, &HealthServiceServer{})

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

var _ healthv1.HealthServiceServer = (*HealthServiceServer)(nil)

type HealthServiceServer struct{}

func (s *HealthServiceServer) Check(ctx context.Context, req *healthv1.CheckRequest) (*healthv1.CheckResponse, error) {
	return &healthv1.CheckResponse{
		Status: healthv1.ServingStatus_SERVING_STATUS_OK,
	}, nil
}

func (s *HealthServiceServer) Watch(req *healthv1.WatchRequest, client healthv1.HealthService_WatchServer) error {
	if err := client.Send(&healthv1.WatchResponse{
		Status: healthv1.ServingStatus_SERVING_STATUS_OK,
	}); err != nil {
		return err
	}
	return nil
}
