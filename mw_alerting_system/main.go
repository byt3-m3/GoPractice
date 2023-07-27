package main

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"log"
)

var (
	manager        = models.NewAlertManager()
	serviceDetails = models.NewServiceDetails(&models.NewServiceDetailsInput{
		Name:    "test-app-beholder",
		Env:     "staging",
		Regions: []string{"us-central1", "us-east4"},
		PromAlerts: []*models.Alert{
			models.NewPrometheusAlert(models.WithTestData()),
		},
		GCPAlerts: []*models.Alert{
			models.NewGCPAlert(models.WithTestData()),
		},
		OpsGenieAlerts: []*models.Alert{
			models.NewOpsGenieAlert(models.WithTestData()),
		},
	})
)

func handlerExample() {

	err := manager.Deploy(context.Background(), serviceDetails,
		handlers.PrometheusAlertDeployHandler,
		handlers.GCPAlertDeployHandler,
		handlers.OpsGenieAlertDeployHandler,
		handlers.MetricHandler,
		handlers.LogInvokeHandler,
	)

	if err != nil {
		log.Println(err)
	}
}

func middlewareExample() {

	if len(serviceDetails.PromAlerts) > 0 {
		err := manager.DeployV2(context.Background(), serviceDetails,
			handlers.PrometheusAlertDeployHandler,
			models.LogMiddleWare,
			models.MetricMiddleWare,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if len(serviceDetails.GCPAlerts) > 0 {
		err := manager.DeployV2(context.Background(), serviceDetails,
			handlers.GCPAlertDeployHandler,
			models.LogMiddleWare,
			models.MetricMiddleWare,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if len(serviceDetails.OpsGenieAlerts) > 0 {
		err := manager.DeployV2(context.Background(), serviceDetails,
			handlers.OpsGenieAlertDeployHandler,
			models.LogMiddleWare,
			models.MetricMiddleWare,
		)
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func main() {
	//middlewareExample()
	handlerExample()
}
