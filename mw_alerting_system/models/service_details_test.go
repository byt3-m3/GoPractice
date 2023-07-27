package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServiceDetails(t *testing.T) {
	t.Run("test when create is successful", func(t *testing.T) {
		sd := NewServiceDetails(&NewServiceDetailsInput{
			Name:           "",
			Env:            "",
			Regions:        nil,
			GCPAlerts:      nil,
			PromAlerts:     nil,
			OpsGenieAlerts: nil,
		})

		assert.IsType(t, &ServiceDetails{}, sd)
	})

	t.Run("test when setting deployContext is successful", func(t *testing.T) {
		sd := NewServiceDetails(&NewServiceDetailsInput{
			Name:           "",
			Env:            "",
			Regions:        nil,
			GCPAlerts:      nil,
			PromAlerts:     nil,
			OpsGenieAlerts: nil,
		})

		sd.SetSourceContext(AlertTypePrometheus, DeployContext{})

		assert.IsType(t, &ServiceDetails{}, sd)
	})

	t.Run("test when getting deployContext is successful", func(t *testing.T) {
		sd := NewServiceDetails(&NewServiceDetailsInput{
			Name:           "",
			Env:            "",
			Regions:        nil,
			GCPAlerts:      nil,
			PromAlerts:     nil,
			OpsGenieAlerts: nil,
		})

		sd.SetSourceContext(AlertTypePrometheus, DeployContext{})

		deployCtx := sd.GetSourceContext(AlertTypePrometheus)

		assert.IsType(t, &DeployContext{}, deployCtx)
	})

	t.Run("test when getting deployContext fails", func(t *testing.T) {
		sd := NewServiceDetails(&NewServiceDetailsInput{
			Name:           "",
			Env:            "",
			Regions:        nil,
			GCPAlerts:      nil,
			PromAlerts:     nil,
			OpsGenieAlerts: nil,
		})

		deployCtx := sd.GetSourceContext(AlertTypePrometheus)

		assert.Nil(t, deployCtx)
	})

	t.Run("test when getting deployContextMap is successful", func(t *testing.T) {
		sd := NewServiceDetails(&NewServiceDetailsInput{
			Name:           "",
			Env:            "",
			Regions:        nil,
			GCPAlerts:      nil,
			PromAlerts:     nil,
			OpsGenieAlerts: nil,
		})

		sd.SetSourceContext(AlertTypePrometheus, DeployContext{})

		deployCtxMap := sd.GetDeployContextMap()

		assert.IsType(t, map[AlertType]DeployContext{}, deployCtxMap)
	})

}
