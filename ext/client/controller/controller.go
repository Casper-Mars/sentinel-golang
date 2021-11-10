package controller

import (
	"context"
	"fmt"
	"github.com/alibaba/sentinel-golang/ext/client/transport"
	"log"
	"net/http"
)

func Start(ctx context.Context, conf *transport.Config) {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
		if err != nil {
			log.Println(err)
		}
	}()
}
