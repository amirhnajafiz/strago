package metrics

import "github.com/prometheus/client_golang/prometheus"

// Metrics
// holds the metrics of strago server.
type Metrics struct {
	numberOfRequests           prometheus.Counter
	numberOfFailedRequests     prometheus.Counter
	numberOfRequestsPerService *prometheus.CounterVec
	responseTime               prometheus.Gauge
}

const (
	namespace = "strago"
	subsystem = "strago"
)

// NewMetrics
// returns a metrics struct to handle metrics of strago server.
func NewMetrics() Metrics {
	return Metrics{
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

// IncRequest
// increase number of requests.
func (m *Metrics) IncRequest() {
	m.numberOfRequests.Inc()
}

// IncRequestPer
// increase number of requests for a service.
func (m *Metrics) IncRequestPer(ip string) {
	m.numberOfRequestsPerService.With(prometheus.Labels{"ip": ip}).Inc()
}

// IncFailed
// increase number of failed requests.
func (m *Metrics) IncFailed() {
	m.numberOfFailedRequests.Inc()
}

// AddResponse
// set a new response time.
func (m *Metrics) AddResponse(duTime float64) {
	m.responseTime.Set(duTime)
}
