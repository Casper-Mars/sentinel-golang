package controller

import (
	"encoding/json"
	"fmt"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/ext/client/model/rule"
	"io/ioutil"
	"net/http"
)

const (
	degradeType = "degrade"
)

func init() {
	http.HandleFunc("/getRules", func(writer http.ResponseWriter, request *http.Request) {
		ruleType := request.URL.Query().Get("type")
		switch ruleType {
		case degradeType:
			degradeHandler(writer, request)
		}
	})
	http.HandleFunc("/setRules", func(writer http.ResponseWriter, request *http.Request) {
		all, err := ioutil.ReadAll(request.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(all))
	})
}

func degradeHandler(writer http.ResponseWriter, request *http.Request) {
	rules := circuitbreaker.GetRules()
	result := make([]rule.Degrade, len(rules))
	for i, item := range rules {
		result[i] = rule.Degrade{
			Resource:           item.Resource,
			Grade:              int32(item.Strategy),
			Count:              item.Threshold,
			TimeWindow:         int32(item.RetryTimeoutMs / 1000),
			MinRequestAmount:   int32(item.MinRequestAmount),
			SlowRatioThreshold: float64(item.MaxAllowedRtMs),
			StatIntervalMs:     int32(item.StatIntervalMs),
			LimitApp:           "default",
		}
	}
	data, err := json.Marshal(result)
	if err != nil {

	}
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(data)
	if err != nil {

	}
}
