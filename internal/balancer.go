package internal

import "sort"

// getOneIPFromServices
// returns one ip address from given services.
// sorting type is based on number of requests, or
// the busy time of a service.
func (s *server) getOneIPFromServices() *service {
	defer func() {
		sort.Slice(s.services, func(i, j int) bool {
			switch s.balancingType {
			case 1:
				return s.services[i].used < s.services[j].used
			case 2:
				return s.services[i].busy < s.services[j].busy
			default:
				return i < j
			}
		})
	}()

	for index := range s.services {
		s.services[index].used++

		return s.services[index]
	}

	return nil
}
