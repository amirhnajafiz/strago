package internal

import "sort"

func (s *server) getIP() string {
	serv := s.services[0]

	serv.used++

	sort.Slice(s.services, func(i, j int) bool {
		return s.services[i].used < s.services[j].used
	})

	return serv.ip
}
