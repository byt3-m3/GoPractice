package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
)

type DeployHandlerInput struct {
	ServiceDetails *models.ServiceDetails
}

type DeployHandler func(ctx context.Context, input DeployHandlerInput) error
