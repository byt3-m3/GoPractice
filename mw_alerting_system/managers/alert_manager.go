package managers

import (
	"context"
	handlers2 "github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/middleware"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
)

type AlertManager interface {
	// Deploy utilizes a simple slice of handlers that are expected to operate on the ServiceDetails
	Deploy(ctx context.Context, sd *models.ServiceDetails, handlers ...handlers2.DeployHandler) error

	// DeployV2 utilizes middleware on a before a single handler is executed
	DeployV2(ctx context.Context, sd *models.ServiceDetails, handler handlers2.DeployHandler, middlewares ...middleware.Middleware) error

	// DeployV3 utilizes middleware on a before a single handler is executed and executes PostProcessing
	DeployV3(ctx context.Context, sd *models.ServiceDetails, handler handlers2.DeployHandler, middlewares []middleware.Middleware, postProcessors []post_processors.PostProcessor) error
}

type alertManager struct {
}

func NewAlertManager() AlertManager {
	return alertManager{}
}

func (a alertManager) Deploy(ctx context.Context, sd *models.ServiceDetails, handlers ...handlers2.DeployHandler) error {

	for _, h := range handlers {
		if err := h(ctx, handlers2.DeployHandlerInput{ServiceDetails: sd}); err != nil {
			return err
		}
	}

	return nil
}

func (a alertManager) DeployV2(ctx context.Context, sd *models.ServiceDetails, handler handlers2.DeployHandler, middlewares ...middleware.Middleware) error {
	for _, mw := range middleware.ReverseMiddlewareSlice(middlewares) {
		handler = mw(handler)

	}

	if err := handler(ctx, handlers2.DeployHandlerInput{ServiceDetails: sd}); err != nil {
		return err
	}

	return nil
}

func (a alertManager) DeployV3(ctx context.Context, sd *models.ServiceDetails, handler handlers2.DeployHandler, middlewares []middleware.Middleware, postProcessors []post_processors.PostProcessor) error {

	// execute middleware in reverse order
	for _, mw := range middleware.ReverseMiddlewareSlice(middlewares) {
		handler = mw(handler)

	}

	// execute main handler
	if err := handler(ctx, handlers2.DeployHandlerInput{ServiceDetails: sd}); err != nil {
		return err
	}

	// execute post-processors
	for _, pp := range postProcessors {
		handler = pp(sd, handler)

	}

	return nil
}
