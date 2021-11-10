package controller

import (
	"encoding/json"

	"github.com/alibaba/sentinel-golang/ext/client/model/api"
)

func init() {
	addApi("/api", apiHandler)
}

func apiHandler(req *request) *response {
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
		return internalErr(nil)
	}
	return success(marshal)
}
