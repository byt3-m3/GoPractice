package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"log"
)

var MetricHandler models.DeployHandler = func(ctx context.Context, input models.DeployHandlerInput) error {

	for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
		if deployContext.IsComplete && deployContext.IsSuccess {
			log.Println("metrics success", source)

		} else {
			log.Println("metrics failure", source)

		}
	}

	return nil
}
