package controller

import (
	"encoding/json"
	"github.com/alibaba/sentinel-golang/ext/client/model/api"
	"net/http"
)

func init() {
	http.HandleFunc("/api", apiHandler)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	result := []api.Api{
		{
			Url:  "/getRules",
			Desc: "get all active rules by type, request param: type={ruleType}",
		},
		{
			Url:  "/setRules",
			Desc: "modify the rules, accept param: type={ruleType}&data={ruleJson}",
		},
		{
			Url:  "/api",
			Desc: "get all available command handlers",
		},
	}
	marshal, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(marshal)
	if err != nil {
		//todo: handle error
	}
}
