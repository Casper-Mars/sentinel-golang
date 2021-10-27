package heartbeat

import (
	"time"
)

type SimpleHttpHeartBeatSender struct {

}

func (s SimpleHttpHeartBeatSender) SendHeartbeat() error {
	panic("implement me")
}

func (s SimpleHttpHeartBeatSender) IntervalMs() time.Duration {
	panic("implement me")
}



