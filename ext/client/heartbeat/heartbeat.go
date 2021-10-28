package heartbeat

import (
	"context"
	"github.com/alibaba/sentinel-golang/ext/client/transport"
	"github.com/alibaba/sentinel-golang/logging"
	"time"
)

func StartHeartbeat(ctx context.Context, conf *transport.Config) error {
	sender := NewSimpleHttpHeartbeatSender(conf)
	err := sender.SendHeartbeat()
	if err != nil {
		return err
	}
	go func() {
		ticker := time.NewTicker(sender.IntervalMs())
		for {
			select {
			case <-ticker.C:
				err := sender.SendHeartbeat()
				if err != nil {
					// TODO: add log
					logging.Error(err, "")
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}
