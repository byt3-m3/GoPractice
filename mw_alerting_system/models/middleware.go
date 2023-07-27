package models

import (
	"context"
	"log"
)

type Middleware func(next DeployHandler) DeployHandler

var (
	VerifyDirectoryMiddleware Middleware = func(next DeployHandler) DeployHandler {

		return func(ctx context.Context, input DeployHandlerInput) error {

			log.Println("directory team is valid")

			return next(ctx, input)
		}
	}

	ValidateInfraspecMiddleware Middleware = func(next DeployHandler) DeployHandler {

		return func(ctx context.Context, input DeployHandlerInput) error {
			log.Println("infrapsec is valid")
			return next(ctx, input)
		}
	}
)

func ReverseMiddlewareSlice(middlewares []Middleware) []Middleware {
	var reveredMiddlewares []Middleware
	for i := len(middlewares) - 1; i >= 0; i-- {

		reveredMiddlewares = append(reveredMiddlewares, middlewares[i])
	}

	return reveredMiddlewares
}

func ReversePostProcessorSlice(middlewares []PostProcessor) []PostProcessor {
	var reveredMiddlewares []PostProcessor
	for i := len(middlewares) - 1; i >= 0; i-- {

		reveredMiddlewares = append(reveredMiddlewares, middlewares[i])
	}

	return reveredMiddlewares
}
