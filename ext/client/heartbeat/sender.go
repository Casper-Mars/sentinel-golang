package heartbeat

import "time"

type Sender interface {
	SendHeartbeat() error
	IntervalMs() time.Duration
}





