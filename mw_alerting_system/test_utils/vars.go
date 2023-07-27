package test_utils

import "github.com/byt3-m3/GoPractice/mw_alerting_system/models"

var (
	TestServiceName = "test-app-beholder"
	TestEnvStaging  = "staging"

	TestRegions = []string{"us-central1", "us-east4"}

	TestGCPAlerts = []*models.Alert{
		models.NewGCPAlert(models.WithTestData()),
		models.NewGCPAlert(models.WithTestData()),
	}

	TestPromAlerts = []*models.Alert{
		models.NewPrometheusAlert(models.WithTestData()),
		models.NewPrometheusAlert(models.WithTestData()),
	}

	TestOpsGenieAlerts = []*models.Alert{
		models.NewOpsGenieAlert(models.WithTestData()),
		models.NewOpsGenieAlert(models.WithTestData()),
	}
)
