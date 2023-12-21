package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"time"

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

	healthv1.RegisterHealthCheckServiceServer(s, &HealthCheckServer{})

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

type HealthCheckServer struct {
	healthv1.UnimplementedHealthCheckServiceServer
}

func (s *HealthCheckServer) Call(ctx context.Context, req *healthv1.CheckRequest) (*healthv1.CheckResponse, error) {
	return &healthv1.CheckResponse{
		Status:  healthv1.ServingStatus_SERVING_STATUS_SERVING,
		Message: req.Message,
	}, nil
}

func (s *HealthCheckServer) Watch(req *healthv1.WatchRequest, watch healthv1.HealthCheckService_WatchServer) error {
	for i := 0; i < int(req.Seconds); i++ {
		if err := watch.Send(&healthv1.WatchResponse{
			Status: healthv1.ServingStatus_SERVING_STATUS_SERVING,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}
