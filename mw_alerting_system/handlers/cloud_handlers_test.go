package handlers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrometheusAlertDeployHandler(t *testing.T) {

	type args struct {
		ctx   context.Context
		input DeployHandlerInput
	}

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "test when handler is successful",
			args: args{
				ctx: context.Background(),
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

			err := PrometheusAlertDeployHandler(tc.args.ctx, tc.args.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}

func TestGCPAlertDeployHandler(t *testing.T) {

	type args struct {
		ctx   context.Context
		input DeployHandlerInput
	}

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "test when handler is successful",
			args: args{
				ctx: context.Background(),
				input: DeployHandlerInput{
					ServiceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
						Name:    test_utils.TestServiceName,
						Env:     test_utils.TestEnvStaging,
						Regions: test_utils.TestRegions,
						GCPAlerts: []*models.Alert{
							models.NewGCPAlert(models.WithTestData()),
						},
					}),
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := GCPAlertDeployHandler(tc.args.ctx, tc.args.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}

func TestOpsGenieAlertDeployHandler(t *testing.T) {

	type args struct {
		ctx   context.Context
		input DeployHandlerInput
	}

	type testCase struct {
		name    string
		args    args
		wantErr bool
	}

	testCases := []testCase{
		{
			name: "test when handler is successful",
			args: args{
				ctx: context.Background(),
				input: DeployHandlerInput{
					ServiceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
						Name:    test_utils.TestServiceName,
						Env:     test_utils.TestEnvStaging,
						Regions: test_utils.TestRegions,
						OpsGenieAlerts: []*models.Alert{
							models.NewGCPAlert(models.WithTestData()),
						},
					}),
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := OpsGenieAlertDeployHandler(tc.args.ctx, tc.args.input)

			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
