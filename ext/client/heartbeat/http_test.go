package heartbeat

import (
	"github.com/alibaba/sentinel-golang/ext/client/transport"
	"testing"
)

func TestSimpleHttpHeartBeatSender_SendHeartbeat(t *testing.T) {
	sender := NewSimpleHttpHeartbeatSender(&transport.Config{
		Server:          "http://localhost:8080",
		Port:            "10086",
		IntervalMs:      2000,
		HeartbeatApi:    "/registry/machine",
		ClientIp:        "127.0.0.1",
		AppType:         "0",
		AppName:         "test-sender",
		Hostname:        "test-sender",
		SentinelVersion: "1.8.0",
	})
	err := sender.SendHeartbeat()
	if err != nil {
		t.Error(err)
	}
}
