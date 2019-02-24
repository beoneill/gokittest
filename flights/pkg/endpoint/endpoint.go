package endpoint

import (
	service "github.com/beoneill/gokittest/flights/pkg/service"
	"context"
	endpoint "github.com/go-kit/kit/endpoint"
)

// HealthRequest collects the request parameters for the Health method.
type HealthRequest struct {
	S string `json:"s"`
}

// HealthResponse collects the response parameters for the Health method.
type HealthResponse struct {
	Rs  string `json:"rs"`
	Err error  `json:"err"`
}

// MakeHealthEndpoint returns an endpoint that invokes Health on the service.
func MakeHealthEndpoint(s service.FlightsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HealthRequest)
		rs, err := s.Health(ctx, req.S)
		return HealthResponse{
			Err: err,
			Rs:  rs,
		}, nil
	}
}

// Failed implements Failer.
func (r HealthResponse) Failed() error {
	return r.Err
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Health implements Service. Primarily useful in a client.
func (e Endpoints) Health(ctx context.Context, s string) (rs string, err error) {
	request := HealthRequest{S: s}
	response, err := e.HealthEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(HealthResponse).Rs, response.(HealthResponse).Err
}
