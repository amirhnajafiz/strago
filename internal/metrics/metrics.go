package metrics

// Metrics
// holds the metrics of strago server.
type Metrics struct {
	numberOfRequests           int
	numberOfFailedRequests     int
	numberOfRequestsPerService map[string]int
	responseTime               []float64
}

// NewMetrics
// returns a metrics struct to handle metrics of strago server.
func NewMetrics() Metrics {
	return Metrics{
		numberOfRequests:           0,
		numberOfFailedRequests:     0,
		numberOfRequestsPerService: make(map[string]int),
	}
}

// IncRequest
// increase number of requests.
func (m *Metrics) IncRequest() {
	m.numberOfRequests++
}

// IncRequestPer
// increase number of requests for a service.
func (m *Metrics) IncRequestPer(ip string) {
	if _, ok := m.numberOfRequestsPerService[ip]; ok {
		m.numberOfRequestsPerService[ip]++
	} else {
		m.numberOfRequestsPerService[ip] = 1
	}
}

// IncFailed
// increase number of failed requests.
func (m *Metrics) IncFailed() {
	m.numberOfFailedRequests++
}

// AddResponse
// set a new response time.
func (m *Metrics) AddResponse(duTime float64) {
	m.responseTime = append(m.responseTime, duTime)
}
