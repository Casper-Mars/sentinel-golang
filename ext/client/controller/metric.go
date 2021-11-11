package controller

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/log/metric"
	"github.com/alibaba/sentinel-golang/logging"
	"strconv"
	"strings"
)

func init() {
	addApi("/metric", metricHandler)
}

func metricHandler(req *request) *response {

	// todo: cache the searcher
	searcher, err := metric.NewDefaultMetricSearcher(config.LogBaseDir(), metric.FormMetricFileName(config.AppName(), false))
	if err != nil {
		logging.Error(err, "fail to create metric searcher")
		return internalErr([]byte(err.Error()))
	}

	startTime := req.URL.Query().Get("startTime")
	endTime := req.URL.Query().Get("endTime")
	maxLines := req.URL.Query().Get("maxLines")
	identity := req.URL.Query().Get("identity")
	if startTime == "" {
		return success([]byte(""))
	}
	startTimeMs, err := strconv.ParseUint(startTime, 10, 64)
	if err != nil {
		logging.Error(err, fmt.Sprintf("fail to parse startTime[%s]", startTime))
		return internalErr([]byte("parameter [startTime] must be number"))
	}
	var lines []*base.MetricItem
	if endTime == "" {
		var ml uint64
		if maxLines == "" {
			ml, err = strconv.ParseUint(maxLines, 10, 32)
			if err != nil {
				logging.Error(err, "fail to parse maxLines")
				return internalErr([]byte("parameter [maxLines] must be number"))
			}
		}
		ml = func(a, b uint64) uint64 {
			if a > b {
				return a
			}
			return b
		}(ml, 12000)
		lines, err = searcher.FindFromTimeWithMaxLines(startTimeMs, uint32(ml))
		if err != nil {
			logging.Error(err, fmt.Sprintf("fail to find metric[startTime:%s,maxLines:%d]", startTime, ml))
			return internalErr([]byte(err.Error()))
		}
	} else {
		endTimeMs, err := strconv.ParseUint(endTime, 10, 64)
		if err != nil {
			logging.Error(err, fmt.Sprintf("fail to parse endTime[%s]", endTime))
			return internalErr([]byte("parameter [endTime] must be number"))
		}
		lines, err = searcher.FindByTimeAndResource(startTimeMs, endTimeMs, identity)
		if err != nil {
			logging.Error(err, fmt.Sprintf("fail to find metric[startTime:%s,endTime:%s,identity:%s]", startTime, endTime, identity))
			return internalErr([]byte(err.Error()))
		}
	}
	sb := strings.Builder{}
	for _, line := range lines {
		thinString, err := line.ToThinString()
		if err != nil {
			logging.Error(err, "fail to get metric item thin string")
			return internalErr([]byte(err.Error()))
		}
		sb.WriteString(thinString)
		sb.WriteString("\n")
	}
	return success([]byte(sb.String()))
}
