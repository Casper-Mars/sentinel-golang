package client

import (
	"context"
	"github.com/alibaba/sentinel-golang/ext/client/config"
	"github.com/alibaba/sentinel-golang/ext/client/controller"
	"github.com/alibaba/sentinel-golang/ext/client/heartbeat"
)

func Init(ctx context.Context, conf *config.Config) error {

	err := heartbeat.Start(ctx, conf)
	if err != nil {
		return err
	}
	controller.Start(ctx, conf)
	return nil
}
