package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewGCPAlert(t *testing.T) {

	alert := NewGCPAlert(WithID("1"), WithAlertName("test-gcp-alert-1"), WithAlertExpr("match_this"))
	alertB := NewGCPAlert(WithTestData())

	assert.Equal(t, AlertTypeGCP, alert.AlertType)
	assert.Equal(t, AlertTypeGCP, alertB.AlertType)

}

func TestNewPrometheusAlert(t *testing.T) {

	alert := NewPrometheusAlert(WithID("1"), WithAlertName("test-prom-alert-1"), WithAlertExpr("match_this"))
	alertB := NewPrometheusAlert(WithTestData())

	assert.Equal(t, AlertTypePrometheus, alert.AlertType)
	assert.Equal(t, AlertTypePrometheus, alertB.AlertType)

}

func TestNewOpsGenieAlert(t *testing.T) {

	alert := NewOpsGenieAlert(WithID("1"), WithAlertName("test-ops-alert-1"), WithAlertExpr("match_this"))
	alertB := NewOpsGenieAlert(WithTestData())

	assert.Equal(t, AlertTypeOpsGenie, alert.AlertType)
	assert.Equal(t, AlertTypeOpsGenie, alertB.AlertType)

}
