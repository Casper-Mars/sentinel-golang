package controller

import (
	"context"
	"fmt"
	"github.com/alibaba/sentinel-golang/ext/client/transport"
	"log"
	"net/http"
)

func Start(ctx context.Context, conf *transport.Config) {
	mux := http.NewServeMux()
	registryRuleController(mux)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), mux)
		if err != nil {
			log.Println(err)
		}
	}()
}
