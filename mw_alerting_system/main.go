package main

import (
	"context"
	"fmt"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/managers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/middleware"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
	"log"
)

var (
	manager         = managers.NewAlertManager()
	serviceDetailsA = models.NewServiceDetails(&models.NewServiceDetailsInput{
		Name:    "test-app-beholder-a",
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
	serviceDetailsB = models.NewServiceDetails(&models.NewServiceDetailsInput{
		Name:    "test-app-beholder-b",
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
	serviceDetailsC = models.NewServiceDetails(&models.NewServiceDetailsInput{
		Name:    "test-app-beholder-c",
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

func SetUpAlertsWithHandlers(sd *models.ServiceDetails, alertTypes ...models.AlertType) {

	for _, alertType := range alertTypes {

		switch alertType {
		case models.AlertTypePrometheus:
			err := manager.Deploy(context.Background(), sd,
				handlers.PrometheusAlertDeployHandler,
				handlers.MetricHandler,
				handlers.LogHandler,
			)
			if err != nil {
				log.Println(err)
			}

		case models.AlertTypeGCP:
			err := manager.Deploy(context.Background(), sd,
				handlers.GCPAlertDeployHandler,
				handlers.MetricHandler,
				handlers.LogHandler,
			)
			if err != nil {
				log.Println(err)
			}

		case models.AlertTypeOpsGenie:
			err := manager.Deploy(context.Background(), sd,
				handlers.OpsGenieAlertDeployHandler,
				handlers.MetricHandler,
				handlers.LogHandler,
			)
			if err != nil {
				log.Println(err)
			}

		case models.AlertTypeAll:
			err := manager.Deploy(context.Background(), sd,
				handlers.PrometheusAlertDeployHandler,
				handlers.GCPAlertDeployHandler,
				handlers.OpsGenieAlertDeployHandler,
				handlers.MetricHandler,
				handlers.LogHandler,
			)
			if err != nil {
				log.Println(err)
			}

		}
	}

}

func SetUpAlertsWithMiddleware(sd *models.ServiceDetails, alertTypes ...models.AlertType) {

	for _, alertType := range alertTypes {

		switch alertType {
		case models.AlertTypePrometheus:
			err := manager.DeployV2(context.Background(), sd,
				handlers.PrometheusAlertDeployHandler,
				middleware.VerifyDirectoryMiddleware,
				middleware.ValidateInfraspecMiddleware,
			)
			if err != nil {
				log.Fatalln(err)
			}

		case models.AlertTypeGCP:
			err := manager.DeployV2(context.Background(), sd,
				handlers.GCPAlertDeployHandler,
				middleware.VerifyDirectoryMiddleware,
				middleware.ValidateInfraspecMiddleware,
			)
			if err != nil {
				log.Fatalln(err)
			}

		case models.AlertTypeOpsGenie:
			err := manager.DeployV2(context.Background(), sd,
				handlers.OpsGenieAlertDeployHandler,
				middleware.VerifyDirectoryMiddleware,
				middleware.ValidateInfraspecMiddleware,
			)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

}

func SetUpAlertsWithMiddlewarePostProcessing(sd *models.ServiceDetails, alertTypes ...models.AlertType) {

	for _, alertType := range alertTypes {

		switch alertType {

		case models.AlertTypePrometheus:
			err := manager.DeployV3(context.Background(), sd,
				handlers.PrometheusAlertDeployHandler,
				[]middleware.Middleware{
					middleware.VerifyDirectoryMiddleware,
					middleware.ValidateInfraspecMiddleware,
				},
				[]post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
					post_processors.MetricPostProcessor,
					post_processors.VerifySourceDeploy,
				},
			)
			if err != nil {
				log.Fatalln(err)
			}

		case models.AlertTypeGCP:
			err := manager.DeployV3(context.Background(), sd,
				handlers.GCPAlertDeployHandler,
				[]middleware.Middleware{
					middleware.VerifyDirectoryMiddleware,
					middleware.ValidateInfraspecMiddleware,
				},
				[]post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
					post_processors.MetricPostProcessor,
					post_processors.VerifySourceDeploy,
				},
			)
			if err != nil {
				log.Fatalln(err)
			}

		case models.AlertTypeOpsGenie:
			err := manager.DeployV3(context.Background(), sd,
				handlers.OpsGenieAlertDeployHandler,
				[]middleware.Middleware{
					middleware.VerifyDirectoryMiddleware,
					middleware.ValidateInfraspecMiddleware,
				},
				[]post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
					post_processors.MetricPostProcessor,
					post_processors.VerifySourceDeploy,
				},
			)
			if err != nil {
				log.Fatalln(err)
			}

		}

	}

}

func main() {
	fmt.Println("\nSetUpAlertsWithHandlers")
	SetUpAlertsWithHandlers(serviceDetailsA, models.AlertTypeAll)

	fmt.Println("\nSetUpAlertsWithMiddleware")
	SetUpAlertsWithMiddleware(serviceDetailsB, models.AlertTypePrometheus)

	fmt.Println("\nSetUpAlertsWithMiddlewarePostProcessing")
	SetUpAlertsWithMiddlewarePostProcessing(serviceDetailsC, models.AlertTypeGCP)
}
