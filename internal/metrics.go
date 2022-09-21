package internal

import "github.com/prometheus/client_golang/prometheus"

// metrics
// holds the metrics of strago server.
type metrics struct {
	numberOfRequests           prometheus.Counter
	numberOfFailedRequests     prometheus.Counter
	numberOfRequestsPerService *prometheus.CounterVec
	responseTime               prometheus.Gauge
}

const (
	namespace = "strago"
	subsystem = "monitoring"
)

// newMetrics
// returns a metrics struct to handle metrics of strago server.
func newMetrics() metrics {
	return metrics{
		numberOfRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "number_of_requests",
			Help:      "total number of requests that are sent to strago server",
		}),
		numberOfFailedRequests: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "number_of_failed_requests",
			Help:      "number of requests that are failed to process",
		}),
		numberOfRequestsPerService: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "number_of_requests_per_service",
			Help:      "total number of requests for each service",
		}, []string{
			"ip",
		}),
		responseTime: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "response_time",
			Help:      "response time of strago service",
		}),
	}
}

// incRequest
// increase number of requests.
func (m *metrics) incRequest() {
	m.numberOfRequests.Inc()
}

// incRequestPer
// increase number of requests for a service.
func (m *metrics) incRequestPer(ip string) {
	m.numberOfRequestsPerService.With(prometheus.Labels{"ip": ip}).Inc()
}

// incFailed
// increase number of failed requests.
func (m *metrics) incFailed() {
	m.numberOfFailedRequests.Inc()
}

// addResponse
// set a new response time.
func (m *metrics) addResponse(duTime float64) {
	m.responseTime.Set(duTime)
}
