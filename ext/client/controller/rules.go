package controller

import (
	"encoding/json"
	"net/url"

	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/ext/client/model/rule"
)

const (
	degradeType = "degrade"
)

func init() {
	addApi("/getRules", func(req *request) *response {
		ruleType := req.URL.Query().Get("type")
		switch ruleType {
		case degradeType:
			return getDegradeRulesHandler(req)
		}
		return success(nil)
	})
	addApi("/setRules", func(req *request) *response {
		ruleType := req.PostForm.Get("type")
		switch ruleType {
		case degradeType:
			setDegradeRulesHandler(req)
		}
		return success(nil)
	})
}

// getDegradeRulesHandler returns all degrade rules
func getDegradeRulesHandler(req *request) *response {
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
	return success(data)
}

// setDegradeRulesHandler sets degrade rules
func setDegradeRulesHandler(req *request) *response {
	data := req.PostForm.Get("data")
	data, err := url.QueryUnescape(data)
	if err != nil {
		// todo: handle error
		return internalErr(nil)
	}
	var rules []rule.Degrade
	err = json.Unmarshal([]byte(data), &rules)
	if err != nil {
		// todo: handle error
		return internalErr(nil)
	}
	newRules := make([]*circuitbreaker.Rule, len(rules))
	for i, item := range rules {
		newRules[i] = &circuitbreaker.Rule{
			Resource:         item.Resource,
			Strategy:         circuitbreaker.Strategy(item.Grade),
			Threshold:        item.Count,
			RetryTimeoutMs:   uint32(item.TimeWindow * 1000),
			MinRequestAmount: uint64(item.MinRequestAmount),
			MaxAllowedRtMs:   uint64(item.SlowRatioThreshold * 1000),
			StatIntervalMs:   uint32(item.StatIntervalMs),
		}
	}
	_, err = circuitbreaker.LoadRules(newRules)
	if err != nil {
		// todo: handle error
		return internalErr(nil)
	}
	return success([]byte("success"))
}
