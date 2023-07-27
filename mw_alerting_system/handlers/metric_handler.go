package handlers

import (
	"context"
	"log"
)

var MetricHandler DeployHandler = func(ctx context.Context, input DeployHandlerInput) error {

	for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
		if deployContext.IsComplete && deployContext.IsSuccess {
			log.Println("metrics success", source)

		} else {
			log.Println("metrics failure", source)

		}
	}

	return nil
}
