package rule

type Degrade struct {
	Resource           string  `json:"resource"`
	Grade              int32   `json:"grade"`
	Count              float64 `json:"count"`
	TimeWindow         int32   `json:"timeWindow"`
	MinRequestAmount   int32   `json:"minRequestAmount"`
	SlowRatioThreshold float64 `json:"slowRatioThreshold"`
	StatIntervalMs     int32   `json:"statIntervalMs"`
	LimitApp           string  `json:"limitApp"`
}
