package managers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/middleware"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlertManager_DeployV3(t *testing.T) {

	manager := NewAlertManager()

	t.Run("test deployV2", func(t *testing.T) {

		sd := models.NewServiceDetails(&models.NewServiceDetailsInput{})

		err := manager.DeployV3(context.Background(), sd, handlers.PrometheusAlertDeployHandler,
			[]middleware.Middleware{
				middleware.ValidateInfraspecMiddleware,
				middleware.VerifyDirectoryMiddleware,
			},
			[]post_processors.PostProcessor{
				post_processors.LoggingPostProcessor,
			},
		)

		assert.NoError(t, err)

	})
}
func TestAlertManager_DeployV2(t *testing.T) {

	manager := NewAlertManager()

	t.Run("test deployV2", func(t *testing.T) {

		sd := models.NewServiceDetails(&models.NewServiceDetailsInput{})

		err := manager.DeployV2(context.Background(), sd, handlers.PrometheusAlertDeployHandler,
			middleware.ValidateInfraspecMiddleware,
			middleware.VerifyDirectoryMiddleware,
		)

		assert.NoError(t, err)

	})
}

func TestAlertManager_Deploy(t *testing.T) {

	manager := NewAlertManager()

	t.Run("test deployV2", func(t *testing.T) {

		sd := models.NewServiceDetails(&models.NewServiceDetailsInput{})

		err := manager.Deploy(context.Background(), sd, handlers.PrometheusAlertDeployHandler)

		assert.NoError(t, err)

	})
}
