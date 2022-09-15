package internal

import "sort"

// getOneIPFromServices
// returns one ip address from given services.
// sorting type is based on number of requests, or
// the busy time of a service.
func (s *server) getOneIPFromServices() string {
	serv := s.services[0]

	serv.used++

	sort.Slice(s.services, func(i, j int) bool {
		return s.services[i].used < s.services[j].used
	})

	return serv.ip
}
