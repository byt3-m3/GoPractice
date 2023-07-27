package models

import (
	"context"
	"log"
)

type Middleware func(next DeployHandler) DeployHandler

var (
	LogMiddleWare Middleware = func(next DeployHandler) DeployHandler {

		return func(ctx context.Context, input DeployHandlerInput) error {
			log.Println("log middleware invoked:", input.ServiceDetails)
			for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
				if deployContext.IsComplete && deployContext.IsSuccess {
					log.Println("successfully deployed:", source)

				} else {
					log.Println("Failed deployed", source)

				}
			}
			return next(ctx, input)
		}

	}

	MetricMiddleWare Middleware = func(next DeployHandler) DeployHandler {
		return func(ctx context.Context, input DeployHandlerInput) error {
			log.Println("log metric invoked:", input.ServiceDetails)

			for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
				if deployContext.IsComplete && deployContext.IsSuccess {
					log.Println("metrics success:", source)

				} else {
					log.Println("metrics failure", source)

				}
			}
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
