package heartbeat

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/ext/client/transport"
	"net/http"
	"strings"
	"time"
)

type SimpleHttpHeartBeatSender struct {
	config *transport.Config
	msg    Message
}

func NewSimpleHttpHeartbeatSender(config *transport.Config) *SimpleHttpHeartBeatSender {
	return &SimpleHttpHeartBeatSender{
		config: config,
		msg: NewMessage(
			config.Port,
			WithIp(config.ClientIp),
			WithHostname(config.Hostname),
			WithApp(config.AppName),
			WithAppType(config.AppType),
			WithSentinelVersion(config.SentinelVersion),
		),
	}
}

func (s SimpleHttpHeartBeatSender) SendHeartbeat() error {
	client := &http.Client{
		Timeout: time.Second * 2,
	}
	sb := strings.Builder{}
	for key, value := range s.msg {
		sb.WriteString(fmt.Sprintf("%s=%s&", key, value))
	}
	param := ""
	if sb.Len() > 0 {
		param = sb.String()[:sb.Len()-1]
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", s.config.Server, s.config.HeartbeatApi), strings.NewReader(param))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (s SimpleHttpHeartBeatSender) IntervalMs() time.Duration {
	return time.Duration(s.config.IntervalMs)
}
