package service

import "context"

// FlightsService describes the service.
type FlightsService interface {
	// Add your methods here

	Health(ctx context.Context, s string) (rs string, err error)
}

type basicFlightsService struct{}

func (b *basicFlightsService) Health(ctx context.Context, s string) (rs string, err error) {
	// TODO implement the business logic of Health

	ok := "OK"

	return ok, err
}

// NewBasicFlightsService returns a naive, stateless implementation of FlightsService.
func NewBasicFlightsService() FlightsService {
	return &basicFlightsService{}
}

// New returns a FlightsService with all of the expected middleware wired in.
func New(middleware []Middleware) FlightsService {
	var svc FlightsService = NewBasicFlightsService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
