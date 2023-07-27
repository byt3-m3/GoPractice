package app

import (
	"context"
	"errors"
	kafka2 "github.com/byt3-m3/GoPractice/mw_alerting_system/kafka"
	"github.com/segmentio/kafka-go"
	"log"
	"testing"
)

func TestApp(t *testing.T) {
	executions := 0

	//liveTestApp := NewApp(WithAlertManager(), WithLiveMessageReader("test-topic", "localhost", "9092"), WithDeployVersion(3))

	messageReaderMock := kafka2.MockMessageReader{
		ReadMockReturn: func() kafka2.ReadMockReturn {
			log.Println("reading mock message")
			executions += 1
			if executions > 1 {
				return kafka2.ReadMockReturn{
					Err: errors.New("finished"),
				}
			}

			return kafka2.ReadMockReturn{
				Message: kafka.Message{
					Topic:         "test-topic",
					Partition:     0,
					Offset:        0,
					HighWaterMark: 0,
					Key:           nil,
					Value:         []byte("{\n\t\"Name\": \"test-app-beholder\",\n\t\"Env\": \"staging\",\n\t\"Regions\": [\n\t\t\"us-central1\",\n\t\t\"us-east4\"\n\t],\n\t\"PromAlerts\": [\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"prometheus\"\n\t\t},\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"prometheus\"\n\t\t}\n\t],\n\t\"GCPAlerts\": [\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"gcp\"\n\t\t},\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"gcp\"\n\t\t}\n\t],\n\t\"OpsGenieAlerts\": [\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"ops_genie\"\n\t\t},\n\t\t{\n\t\t\t\"ID\": \"test-ID\",\n\t\t\t\"Name\": \"test-alert\",\n\t\t\t\"Expr\": \"test-Expr\",\n\t\t\t\"AlertType\": \"ops_genie\"\n\t\t}\n\t]\n}"),
					Headers:       nil,
				},
				Err: nil,
			}
		}}

	testApp := NewApp(WithAlertManager(), WithMockMessageReader(messageReaderMock), WithDeployVersion(3))

	if err := testApp.Run(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
