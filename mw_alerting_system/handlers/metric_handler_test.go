package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetricHandler(t *testing.T) {

	type args struct {
		ctx       context.Context
		input     DeployHandlerInput
		deployCtx models.DeployContext
	}

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "test when deployContext deployed but failed",
			args: args{
				ctx: context.Background(),
				deployCtx: models.DeployContext{
					IsComplete: true,
					IsSuccess:  false,
				},
				input: DeployHandlerInput{
					ServiceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
						Name:    test_utils.TestServiceName,
						Env:     test_utils.TestEnvStaging,
						Regions: test_utils.TestRegions,
						PromAlerts: []*models.Alert{
							models.NewPrometheusAlert(models.WithTestData()),
						},
					}),
				},
			},
			wantErr: false,
		},
		{
			name: "test when deployContext deployed and succeeded",
			args: args{
				ctx: context.Background(),
				deployCtx: models.DeployContext{
					IsComplete: true,
					IsSuccess:  true,
				},
				input: DeployHandlerInput{
					ServiceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
						Name:    test_utils.TestServiceName,
						Env:     test_utils.TestEnvStaging,
						Regions: test_utils.TestRegions,
						PromAlerts: []*models.Alert{
							models.NewPrometheusAlert(models.WithTestData()),
						},
					}),
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tc.args.input.ServiceDetails.SetSourceContext(models.AlertTypeGCP, tc.args.deployCtx)

			err := MetricHandler(tc.args.ctx, tc.args.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
