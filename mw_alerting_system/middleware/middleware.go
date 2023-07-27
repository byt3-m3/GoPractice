package middleware

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
	"log"
)

type Middleware func(next handlers.DeployHandler) handlers.DeployHandler

var (
	VerifyDirectoryMiddleware Middleware = func(next handlers.DeployHandler) handlers.DeployHandler {

		return func(ctx context.Context, input handlers.DeployHandlerInput) error {

			log.Println("directory team is valid")

			return next(ctx, input)
		}
	}

	ValidateInfraspecMiddleware Middleware = func(next handlers.DeployHandler) handlers.DeployHandler {

		return func(ctx context.Context, input handlers.DeployHandlerInput) error {
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

func ReversePostProcessorSlice(middlewares []post_processors.PostProcessor) []post_processors.PostProcessor {
	var reveredMiddlewares []post_processors.PostProcessor
	for i := len(middlewares) - 1; i >= 0; i-- {

		reveredMiddlewares = append(reveredMiddlewares, middlewares[i])
	}

	return reveredMiddlewares
}
