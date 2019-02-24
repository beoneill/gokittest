package grpc

import (
	endpoint "github.com/beoneill/gokittest/flights/pkg/endpoint"
	pb "github.com/beoneill/gokittest/flights/pkg/grpc/pb"
	"context"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeHealthHandler creates the handler logic
func makeHealthHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.HealthEndpoint, decodeHealthRequest, encodeHealthResponse, options...)
}

// decodeHealthResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain sum request.
// TODO implement the decoder
func decodeHealthRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Flights' Decoder is not impelemented")
}

// encodeHealthResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeHealthResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Flights' Encoder is not impelemented")
}
func (g *grpcServer) Health(ctx context1.Context, req *pb.HealthRequest) (*pb.HealthReply, error) {
	_, rep, err := g.health.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.HealthReply), nil
}
