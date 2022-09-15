package internal

import "fmt"

func (s *server) changeStatusForAService(ip string, status bool) error {
	for _, serv := range s.services {
		if serv.ip == ip {
			serv.enable = status

			return nil
		}
	}

	return fmt.Errorf("service with ip '%s' not found", ip)
}
