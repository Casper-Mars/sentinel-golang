package heartbeat

import (
	"strconv"
	"time"
)

type MsgOption func(m Message)

type Message map[string]string

func NewMessage(port string, opts ...MsgOption) Message {
	msg := Message{
		"hostname": "unknown",
		"ip":       "unknown",
		"app":      "unknown",
		"app_type": "0",
		"v":        "unknown",
		"port":     port,
		"version":  strconv.Itoa(int(time.Now().Unix())),
	}
	for _, opt := range opts {
		opt(msg)
	}
	return msg
}

func WithHostname(hostname string) MsgOption {
	return func(m Message) {
		m["hostname"] = hostname
	}
}

func WithIp(ip string) MsgOption {
	return func(m Message) {
		m["ip"] = ip
	}
}

func WithApp(app string) MsgOption {
	return func(m Message) {
		m["app"] = app
	}
}

func WithAppType(appType string) MsgOption {
	return func(m Message) {
		m["app_type"] = appType
	}
}

func WithSentinelVersion(v string) MsgOption {
	return func(m Message) {
		m["v"] = v
	}
}

type Sender interface {
	SendHeartbeat() error
	IntervalMs() time.Duration
}
