package controller

import (
	"net/http"
)

type request struct {
	http.Request
}

func newRequest(r *http.Request) *request {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	return &request{
		Request: *r,
	}
}
