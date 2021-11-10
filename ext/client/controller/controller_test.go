package controller

import (
	"context"
	"testing"

	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/ext/client/heartbeat"
	"github.com/alibaba/sentinel-golang/ext/client/transport"
)

func TestController(t *testing.T) {
	circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		{
			Resource:       "test",
			Threshold:      0.1,
			StatIntervalMs: 1000,
			Strategy:       1,
			RetryTimeoutMs: 1000,
		}, {
			Resource:       "test2",
			Threshold:      0.1,
			StatIntervalMs: 1000,
			Strategy:       1,
			RetryTimeoutMs: 1000,
		},
	})

	config := transport.Config{
		Server:          "http://localhost:8080",
		Port:            "10086",
		IntervalMs:      2000,
		HeartbeatApi:    "/registry/machine",
		ClientIp:        "host.docker.internal",
		AppType:         "0",
		AppName:         "test-sender",
		Hostname:        "test-sender",
		SentinelVersion: "1.8.0",
	}
	Start(context.Background(), &config)
	heartbeat.Start(context.Background(), &config)
	select {}
}
