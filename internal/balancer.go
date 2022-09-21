package internal

import (
	"sort"
)

// getOneIPFromServices
// returns one ip address from given services.
// sorting type is based on number of requests, or
// the busy time of a service.
func (s *server) getOneIPFromServices() *service {
	defer func() {
		sort.Slice(s.services, func(i, j int) bool {
			return s.services[i].used < s.services[j].used
		})
	}()

	for index := range s.services {
		if s.services[index].enable {
			s.services[index].used++

			return s.services[index]
		}
	}

	return nil
}
