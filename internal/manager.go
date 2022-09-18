package internal

import (
	"fmt"
)

// changeStatusForAService
// manages to set status for a service.
func (s *server) changeStatusForAService(ip string, status bool) error {
	for _, serv := range s.services {
		if serv.ip == ip {
			serv.enable = status

			return nil
		}
	}

	return fmt.Errorf("service with ip '%s' not found", ip)
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
