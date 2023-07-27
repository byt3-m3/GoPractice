package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"log"
)

func LogHandler(ctx context.Context, input models.DeployHandlerInput) error {
	for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
		if deployContext.IsComplete && deployContext.IsSuccess {
			log.Println("successfully deployed", source)

		} else {
			log.Println("Failed deployed", source)

		}
	}

	return nil
}
