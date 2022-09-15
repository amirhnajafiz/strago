package internal

import "fmt"

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
