package middlewares

import (
	"context"
	"errors"
	"ticketsbooking/libs/log"
	"time"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"
)

func GrpcLogger(logger *log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			tStart := time.Now()
			method, _ := grpc.Method(ctx)

			resp, err := next(ctx, request)

			logger.Info().Str("method", method).Str("duration", time.Since(tStart).String()).Msg("Done")

			return resp, err
		}
	}
}

func Recoverer(logger *log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					logger.Error().Msgf("Recovered: %v", r)
					err = errors.New("internal server error")
				}
			}()
			return next(ctx, request)
		}
	}
}
