package internal

import "time"

// service
// keeps the data of a service that
// has been given to strago.
type service struct {
	// enable/disable service.
	enable bool
	// service ip.
	ip string
	// number of requests that are sent to this service.
	used int
	// busy time of this service
	busy time.Duration
}

// generateServicesFromGiven
// creates the list of the services.
func generateServicesFromGiven(services []string) []*service {
	list := make([]*service, len(services))

	for index, ip := range services {
		list[index] = &service{
			enable: true,
			ip:     ip,
			used:   0,
			busy:   0,
		}
	}

	return list
}
