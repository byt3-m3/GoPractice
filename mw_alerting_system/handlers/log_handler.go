package handlers

import (
	"context"
	"log"
)

func LogHandler(ctx context.Context, input DeployHandlerInput) error {
	for source, deployContext := range input.ServiceDetails.GetDeployContextMap() {
		if deployContext.IsComplete && deployContext.IsSuccess {
			log.Println("successfully deployed", source)

		}
	}

	return nil
}
