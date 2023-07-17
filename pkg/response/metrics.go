package response

type (
	// ExportMetrics is used in order to send current metrics.
	ExportMetrics struct {
		Requests           int            `json:"requests"`
		FailedRequests     int            `json:"failed_requests"`
		RequestsPerService map[string]int `json:"requests_per_service"`
		ResponseTime       []float64      `json:"response_time"`
	}
)
