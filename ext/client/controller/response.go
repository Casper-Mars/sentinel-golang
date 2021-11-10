package controller

import "net/http"

type response struct {
	status int
	data   []byte
}

func newResponse(status int, data []byte) *response {
	return &response{
		status: status,
		data:   data,
	}
}

func success(data []byte) *response {
	return newResponse(http.StatusOK, data)
}

func internalErr(data []byte) *response {
	return newResponse(http.StatusInternalServerError, data)
}
