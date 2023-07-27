package app

import (
	"context"
	"encoding/json"
	handlers2 "github.com/byt3-m3/GoPractice/mw_alerting_system/handlers"
	kafka2 "github.com/byt3-m3/GoPractice/mw_alerting_system/kafka"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/managers"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/middleware"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/models"
	"github.com/byt3-m3/GoPractice/mw_alerting_system/post_processors"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type AppOpt func(app *app)

var (
	WithLiveMessageReader = func(topic, host, port string) AppOpt {

		return func(app *app) {
			conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, 0)
			if err != nil {
				log.Fatal("failed to dial leader:", err)
			}
			app.KafkaMessageReader = kafka2.NewMessageReader(conn)

		}

	}

	WithMockMessageReader = func(mock kafka2.MockMessageReader) AppOpt {

		return func(app *app) {

			app.KafkaMessageReader = mock
		}
	}
	WithAlertManager = func() AppOpt {
		return func(app *app) {
			app.alertManager = managers.NewAlertManager()
		}
	}

	WithDeployVersion = func(version int) AppOpt {
		return func(app *app) {
			app.deployVersion = version
		}
	}
)

type App interface {
	Run(ctx context.Context) error
}

type Config struct {
}

type app struct {
	alertManager       managers.AlertManager
	kafkaConnection    *kafka.Conn
	KafkaMessageReader kafka2.MessageReader
	deployVersion      int
}

func NewApp(opt ...AppOpt) App {

	a := &app{}

	for _, o := range opt {
		o(a)
	}

	if a.deployVersion == 0 {
		log.Fatal("deploy version is not set use 'WithDeployVersion' ")
	}
	return a
}

func (a *app) Run(ctx context.Context) error {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		log.Println("starting alert manager")

		for {
			m, err := a.KafkaMessageReader.Read()
			if err != nil {
				log.Println("received error from broker")
				return
			}

			sd := models.NewServiceDetails(&models.NewServiceDetailsInput{})
			if err := json.Unmarshal(m.Value, &sd); err != nil {
				log.Fatal(err)
			}

			if sd.HasAlerts(models.AlertTypePrometheus) {
				if a.deployVersion == 1 {
					if err := a.alertManager.Deploy(ctx, sd, handlers2.PrometheusAlertDeployHandler); err != nil {
						log.Fatal(err)
					}

				}

				if a.deployVersion == 2 {
					if err := a.alertManager.DeployV2(ctx, sd, handlers2.PrometheusAlertDeployHandler, middleware.VerifyDirectoryMiddleware); err != nil {
						log.Fatal(err)
					}
				}

				if a.deployVersion == 3 {

					if err := a.alertManager.DeployV3(ctx, sd, handlers2.PrometheusAlertDeployHandler, []middleware.Middleware{
						middleware.ValidateInfraspecMiddleware,
						middleware.VerifyDirectoryMiddleware,
					}, []post_processors.PostProcessor{
						post_processors.LoggingPostProcessor,
						post_processors.MetricPostProcessor,
					}); err != nil {
						log.Fatal(err)
					}

				}

			}

			if sd.HasAlerts(models.AlertTypeGCP) {
				if a.deployVersion == 1 {
					if err := a.alertManager.Deploy(ctx, sd, handlers2.GCPAlertDeployHandler); err != nil {
						log.Fatal(err)
					}

				}

				if a.deployVersion == 2 {
					if err := a.alertManager.DeployV2(ctx, sd, handlers2.GCPAlertDeployHandler, middleware.VerifyDirectoryMiddleware); err != nil {
						log.Fatal(err)
					}
				}

				if a.deployVersion == 3 {

					if err := a.alertManager.DeployV3(ctx, sd, handlers2.GCPAlertDeployHandler, []middleware.Middleware{
						middleware.ValidateInfraspecMiddleware,
						middleware.VerifyDirectoryMiddleware,
					}, []post_processors.PostProcessor{
						post_processors.LoggingPostProcessor,
						post_processors.MetricPostProcessor,
					}); err != nil {
						log.Fatal(err)
					}

				}

			}

			if sd.HasAlerts(models.AlertTypeOpsGenie) {
				if a.deployVersion == 1 {
					if err := a.alertManager.Deploy(ctx, sd, handlers2.OpsGenieAlertDeployHandler); err != nil {
						log.Fatal(err)
					}

				}

				if a.deployVersion == 2 {
					if err := a.alertManager.DeployV2(ctx, sd, handlers2.OpsGenieAlertDeployHandler, middleware.VerifyDirectoryMiddleware); err != nil {
						log.Fatal(err)
					}
				}

				if a.deployVersion == 3 {

					if err := a.alertManager.DeployV3(ctx, sd, handlers2.OpsGenieAlertDeployHandler, []middleware.Middleware{
						middleware.ValidateInfraspecMiddleware,
						middleware.VerifyDirectoryMiddleware,
					}, []post_processors.PostProcessor{
						post_processors.LoggingPostProcessor,
						post_processors.MetricPostProcessor,
					}); err != nil {
						log.Fatal(err)
					}

				}

			}
		}

	}(wg)

	wg.Wait()

	return nil
}
