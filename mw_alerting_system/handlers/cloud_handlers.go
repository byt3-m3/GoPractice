package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"log"
)

func PrometheusAlertDeployHandler(ctx context.Context, input DeployHandlerInput) error {
	log.Println("handling prometheus deployment", input.ServiceDetails.PromAlerts)
	input.ServiceDetails.SetSourceContext(models.AlertTypePrometheus, models.DeployContext{
		IsComplete: true,
		IsSuccess:  true,
	})
	log.Println("handled prometheus deployment")

	return nil
}

func GCPAlertDeployHandler(ctx context.Context, input DeployHandlerInput) error {
	log.Println("handling GCP deployment", input.ServiceDetails.GCPAlerts)

	input.ServiceDetails.SetSourceContext(models.AlertTypeGCP, models.DeployContext{
		IsComplete: true,
		IsSuccess:  true,
	})

	log.Println("handled GCP deployment")

	return nil
}

func OpsGenieAlertDeployHandler(ctx context.Context, input DeployHandlerInput) error {
	log.Println("handling OpsGenie deployment", input.ServiceDetails.OpsGenieAlerts)
	input.ServiceDetails.SetSourceContext(models.AlertTypeOpsGenie, models.DeployContext{
		IsComplete: true,
		IsSuccess:  true,
	})

	log.Println("handled OpsGenie deployment")

	return nil
}
