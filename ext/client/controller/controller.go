package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/alibaba/sentinel-golang/ext/client/transport"
)

type handler func(req *request) *response

func Start(ctx context.Context, conf *transport.Config) {
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)
		if err != nil {
			log.Println(err)
		}
	}()
}

func addApi(pattern string, h handler) {
	http.HandleFunc(pattern, func(writer http.ResponseWriter, r *http.Request) {
		req := newRequest(r)
		resp := h(req)
		writer.WriteHeader(resp.status)
		if resp.data != nil {
			_, err := writer.Write(resp.data)
			if err != nil {
				// todo: handle error
			}
		}
	})
}
