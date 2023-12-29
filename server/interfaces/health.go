package interfaces

import (
	"context"

	healthv1 "github.com/mi11km/neuron-visualizer/server/proto/health/v1"
)

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
