package interfaces

import (
	"context"

	"connectrpc.com/connect"
	healthv1 "github.com/mi11km/neuron-visualizer/server/proto/health/v1"
	"github.com/mi11km/neuron-visualizer/server/proto/health/v1/healthv1connect"
)

var _ healthv1connect.HealthServiceHandler = (*HealthServiceHandler)(nil)

type HealthServiceHandler struct{}

func (h *HealthServiceHandler) Check(
	ctx context.Context, req *connect.Request[healthv1.CheckRequest],
) (*connect.Response[healthv1.CheckResponse], error) {
	return connect.NewResponse(
		&healthv1.CheckResponse{
			Status: healthv1.ServingStatus_SERVING_STATUS_OK,
		},
	), nil
}

func (h *HealthServiceHandler) Watch(
	ctx context.Context, req *connect.Request[healthv1.WatchRequest],
	stream *connect.ServerStream[healthv1.WatchResponse],
) error {
	for i := 0; i < 10; i++ {
		if err := stream.Send(
			&healthv1.WatchResponse{
				Status: healthv1.ServingStatus_SERVING_STATUS_OK,
			},
		); err != nil {
			return err
		}
	}
	return nil
}
