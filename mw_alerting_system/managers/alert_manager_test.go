package managers

import (
	"context"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/middleware"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/test_utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAlertManager_DeployV3(t *testing.T) {

	type args struct {
		ctx            context.Context
		serviceDetails *models.ServiceDetails
		handler        handlers.DeployHandler
		middlewares    []middleware.Middleware
		postProcessors []post_processors.PostProcessor
	}

	type testCase struct {
		name    string
		args    args
		manager AlertManager
		wantErr bool
	}

	testCases := []testCase{

		{
			wantErr: false,
			name:    "test when Prometheus alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     test_utils.TestPromAlerts,
					OpsGenieAlerts: nil,
				}),
				handler: handlers.PrometheusAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
				postProcessors: []post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when GCP alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      test_utils.TestGCPAlerts,
					PromAlerts:     nil,
					OpsGenieAlerts: nil,
				}),
				handler: handlers.GCPAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
				postProcessors: []post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when OpsGenie alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     nil,
					OpsGenieAlerts: test_utils.TestOpsGenieAlerts,
				}),
				handler: handlers.OpsGenieAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
				postProcessors: []post_processors.PostProcessor{
					post_processors.LoggingPostProcessor,
				},
			},
			manager: NewAlertManager(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.manager.DeployV3(tc.args.ctx, tc.args.serviceDetails, tc.args.handler, tc.args.middlewares, tc.args.postProcessors)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
func TestAlertManager_DeployV2(t *testing.T) {

	type args struct {
		ctx            context.Context
		serviceDetails *models.ServiceDetails
		handler        handlers.DeployHandler
		middlewares    []middleware.Middleware
	}

	type testCase struct {
		name    string
		args    args
		manager AlertManager
		wantErr bool
	}

	testCases := []testCase{

		{
			wantErr: false,
			name:    "test when Prometheus alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     test_utils.TestPromAlerts,
					OpsGenieAlerts: nil,
				}),
				handler: handlers.PrometheusAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when GCP alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      test_utils.TestGCPAlerts,
					PromAlerts:     nil,
					OpsGenieAlerts: nil,
				}),
				handler: handlers.GCPAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when OpsGenie alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     nil,
					OpsGenieAlerts: test_utils.TestOpsGenieAlerts,
				}),
				handler: handlers.OpsGenieAlertDeployHandler,
				middlewares: []middleware.Middleware{
					middleware.ValidateInfraspecMiddleware,
					middleware.VerifyDirectoryMiddleware,
				},
			},
			manager: NewAlertManager(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.manager.DeployV2(tc.args.ctx, tc.args.serviceDetails, tc.args.handler, tc.args.middlewares...)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}

func TestAlertManager_Deploy(t *testing.T) {

	type args struct {
		ctx            context.Context
		serviceDetails *models.ServiceDetails

		handlers []handlers.DeployHandler
	}

	type testCase struct {
		name    string
		args    args
		manager AlertManager
		wantErr bool
	}

	testCases := []testCase{

		{
			wantErr: false,
			name:    "test when All alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      test_utils.TestGCPAlerts,
					PromAlerts:     test_utils.TestPromAlerts,
					OpsGenieAlerts: test_utils.TestOpsGenieAlerts,
				}),
				handlers: []handlers.DeployHandler{
					handlers.PrometheusAlertDeployHandler,
					handlers.OpsGenieAlertDeployHandler,
					handlers.GCPAlertDeployHandler,
					handlers.LogHandler,
					handlers.MetricHandler,
				},
			},
			manager: NewAlertManager(),
		},

		{
			wantErr: false,
			name:    "test when Prometheus alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     test_utils.TestPromAlerts,
					OpsGenieAlerts: nil,
				}),
				handlers: []handlers.DeployHandler{
					handlers.PrometheusAlertDeployHandler,
					handlers.LogHandler,
					handlers.MetricHandler,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when GCP alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      test_utils.TestGCPAlerts,
					PromAlerts:     nil,
					OpsGenieAlerts: nil,
				}),
				handlers: []handlers.DeployHandler{
					handlers.GCPAlertDeployHandler,
					handlers.LogHandler,
					handlers.MetricHandler,
				},
			},
			manager: NewAlertManager(),
		},
		{
			wantErr: false,
			name:    "test when OpsGenie alerts deploy is successful",
			args: args{
				ctx: context.Background(),
				serviceDetails: models.NewServiceDetails(&models.NewServiceDetailsInput{
					Name:           test_utils.TestServiceName,
					Env:            test_utils.TestEnvStaging,
					Regions:        test_utils.TestRegions,
					GCPAlerts:      nil,
					PromAlerts:     nil,
					OpsGenieAlerts: test_utils.TestOpsGenieAlerts,
				}),
				handlers: []handlers.DeployHandler{
					handlers.OpsGenieAlertDeployHandler,
					handlers.LogHandler,
					handlers.MetricHandler,
				},
			},
			manager: NewAlertManager(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			err := tc.manager.Deploy(tc.args.ctx, tc.args.serviceDetails, tc.args.handlers...)
			if tc.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

		})
	}

}
