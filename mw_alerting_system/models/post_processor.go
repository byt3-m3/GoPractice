package models

import (
	"context"
	"log"
)

type PostProcessor func(sd *ServiceDetails, next DeployHandler) DeployHandler

var (
	LoggingPostProcessor PostProcessor = func(sd *ServiceDetails, next DeployHandler) DeployHandler {
		deployCtxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range deployCtxMap {
			if serviceCtx.IsComplete && serviceCtx.IsSuccess {
				log.Println("log successfully deployed", service)
			} else {
				log.Println("log failed deployed", service)
			}
		}
		return func(ctx context.Context, input DeployHandlerInput) error {

			return next(ctx, input)
		}
	}

	MetricPostProcessor PostProcessor = func(sd *ServiceDetails, next DeployHandler) DeployHandler {
		deployCtxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range deployCtxMap {
			if serviceCtx.IsComplete && serviceCtx.IsSuccess {
				log.Println("success metric increment", service)
			} else {
				log.Println("failure metric incremented")
			}
		}
		return func(ctx context.Context, input DeployHandlerInput) error {

			return next(ctx, input)
		}
	}

	VerifySourceDeploy PostProcessor = func(sd *ServiceDetails, next DeployHandler) DeployHandler {
		ctxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range ctxMap {

			log.Println(service, "deployment has been verified", serviceCtx)

		}
		return func(ctx context.Context, input DeployHandlerInput) error {
			return next(ctx, input)
		}

	}
)
