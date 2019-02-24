package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(FlightsService) FlightsService

type loggingMiddleware struct {
	logger log.Logger
	next   FlightsService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a FlightsService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next FlightsService) FlightsService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Health(ctx context.Context, s string) (rs string, err error) {
	defer func() {
		l.logger.Log("method", "Health", "s", s, "rs", rs, "err", err)
	}()
	return l.next.Health(ctx, s)
}
