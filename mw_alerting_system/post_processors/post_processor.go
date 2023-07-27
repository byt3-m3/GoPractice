package post_processors

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"log"
)

type PostProcessor func(sd *models.ServiceDetails, next handlers.DeployHandler) handlers.DeployHandler

var (
	LoggingPostProcessor PostProcessor = func(sd *models.ServiceDetails, next handlers.DeployHandler) handlers.DeployHandler {
		deployCtxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range deployCtxMap {
			if serviceCtx.IsComplete && serviceCtx.IsSuccess {
				log.Println("log successfully deployed", service)
			} else {
				log.Println("log failed deployed", service)
			}
		}
		return func(ctx context.Context, input handlers.DeployHandlerInput) error {

			return next(ctx, input)
		}
	}

	MetricPostProcessor PostProcessor = func(sd *models.ServiceDetails, next handlers.DeployHandler) handlers.DeployHandler {
		deployCtxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range deployCtxMap {
			if serviceCtx.IsComplete && serviceCtx.IsSuccess {
				log.Println("success metric increment", service)
			} else {
				log.Println("failure metric incremented")
			}
		}
		return func(ctx context.Context, input handlers.DeployHandlerInput) error {

			return next(ctx, input)
		}
	}

	VerifySourceDeploy PostProcessor = func(sd *models.ServiceDetails, next handlers.DeployHandler) handlers.DeployHandler {
		ctxMap := sd.GetDeployContextMap()

		for service, serviceCtx := range ctxMap {

			log.Println(service, "deployment has been verified", serviceCtx)

		}
		return func(ctx context.Context, input handlers.DeployHandlerInput) error {
			return next(ctx, input)
		}

	}
)
