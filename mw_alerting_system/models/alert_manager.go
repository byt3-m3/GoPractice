package models

import (
	"context"
)

type AlertManager interface {
	// Deploy utilizes a simple slice of handlers that are expected to operate on the ServiceDetails
	Deploy(ctx context.Context, sd *ServiceDetails, handlers ...DeployHandler) error

	// DeployV2 utilizes middleware on a before a single handler is executed
	DeployV2(ctx context.Context, sd *ServiceDetails, handler DeployHandler, middlewares ...Middleware) error

	// DeployV3 utilizes middleware on a before a single handler is executed and executes PostProcessing
	DeployV3(ctx context.Context, sd *ServiceDetails, handler DeployHandler, middlewares []Middleware, postProcessors []PostProcessor) error
}

type alertManager struct {
}

func NewAlertManager() AlertManager {
	return alertManager{}
}

func (a alertManager) Deploy(ctx context.Context, sd *ServiceDetails, handlers ...DeployHandler) error {

	for _, h := range handlers {
		if err := h(ctx, DeployHandlerInput{ServiceDetails: sd}); err != nil {
			return err
		}
	}

	return nil
}

func (a alertManager) DeployV2(ctx context.Context, sd *ServiceDetails, handler DeployHandler, middlewares ...Middleware) error {
	for _, mw := range ReverseMiddlewareSlice(middlewares) {
		handler = mw(handler)

	}

	if err := handler(ctx, DeployHandlerInput{ServiceDetails: sd}); err != nil {
		return err
	}

	return nil
}

func (a alertManager) DeployV3(ctx context.Context, sd *ServiceDetails, handler DeployHandler, middlewares []Middleware, postProcessors []PostProcessor) error {

	// execute middleware in reverse order
	for _, mw := range ReverseMiddlewareSlice(middlewares) {
		handler = mw(handler)

	}

	// execute main handler
	if err := handler(ctx, DeployHandlerInput{ServiceDetails: sd}); err != nil {
		return err
	}

	// execute post-processors
	for _, pp := range postProcessors {
		handler = pp(sd, handler)

	}

	return nil
}
