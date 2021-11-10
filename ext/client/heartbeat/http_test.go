package heartbeat

import (
	"testing"

	"github.com/alibaba/sentinel-golang/ext/client/transport"
)

func TestSimpleHttpHeartBeatSender_SendHeartbeat(t *testing.T) {
	config := transport.Config{
		Server:          "http://localhost:8080",
		Port:            "10086",
		IntervalMs:      2000,
		HeartbeatApi:    "/registry/machine",
		ClientIp:        "127.0.0.1",
		AppType:         "0",
		AppName:         "test-sender",
		Hostname:        "test-sender",
		SentinelVersion: "1.8.0",
	}
	sender := NewSimpleHttpHeartbeatSender(&config, NewMessage(
		config.Port,
		WithApp(config.AppName),
		WithAppType(config.AppType),
		WithHostname(config.Hostname),
		WithSentinelVersion(config.SentinelVersion),
		WithIp(config.ClientIp),
	))
	err := sender.SendHeartbeat()
	if err != nil {
		t.Error(err)
	}
}
