package models

import "context"

type DeployHandlerInput struct {
	ServiceDetails *ServiceDetails
}

type DeployHandler func(ctx context.Context, input DeployHandlerInput) error
