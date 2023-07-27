package models

import (
	"context"
)

type AlertManager interface {
	Deploy(ctx context.Context, sd *ServiceDetails, handlers ...DeployHandler) error
	DeployV2(ctx context.Context, sd *ServiceDetails, handler DeployHandler, middlewares ...Middleware) error
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
	reversedMiddlewares := ReverseMiddlewareSlice(middlewares)
	for _, mw := range reversedMiddlewares {
		handler = mw(handler)

	}

	if err := handler(ctx, DeployHandlerInput{ServiceDetails: sd}); err != nil {
		return err
	}

	return nil
}
