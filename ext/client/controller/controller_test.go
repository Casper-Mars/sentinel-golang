package controller

import (
	"context"
	sentinel "github.com/alibaba/sentinel-golang/api"
	config2 "github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/ext/client/config"
	"github.com/alibaba/sentinel-golang/logging"
	"net/http"
	"testing"

	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/ext/client/heartbeat"
)

func TestController(t *testing.T) {
	defaultConfig := config2.NewDefaultConfig()
	defaultConfig.Sentinel.App.Name = "test-dashboard"
	defaultConfig.Sentinel.Log.Logger = logging.NewConsoleLogger()
	sentinel.InitWithConfig(defaultConfig)
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

	conf := config.Config{
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
	Start(context.Background(), &conf)
	heartbeat.Start(context.Background(), &conf)
	go startTestServer()
	select {}
}

func startTestServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, r *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Add("Content-Type", "application/json")
		entry, blockError := sentinel.Entry("test")
		if blockError != nil {
			writer.Write([]byte(blockError.Error()))
			return
		}
		defer entry.Exit()
		writer.Write([]byte(`{"code":0,"message":"ok"}`))
	})
	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		panic(err)
	}
}
