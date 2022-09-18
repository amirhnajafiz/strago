package internal

// checkIPRangeInBlackList
// goes through blacklist ips and checks
// if the given ip is in the list or not.
func (s *server) checkIPRangeInBlackList(ip string) bool {
	switch s.ipManager.ipType(ip) {
	case ipV6:
		return s.handleIPv6(ip)
	case ipV4:
		return s.handleIPv4(ip)
	}

	return false
}

// handleIPv4
// handles the requests with client ip version 4.
func (s *server) handleIPv4(ip string) bool {
	match := 0

	for _, blackListIP := range s.ipManager.ipV4List {
		ipParts := s.ipManager.getIPv4Parts(ip)
		blackIPParts := s.ipManager.getIPv4Parts(blackListIP)

		for index, part := range ipParts {
			if blackIPParts[index] == "*" || blackIPParts[index] == part {
				match++
			}
		}
	}

	return match == 4
}

// handleIPv6
// handles the requests with client ip version 6.
func (s *server) handleIPv6(ip string) bool {
	match := 0

	for _, blackListIP := range s.ipManager.ipV6List {
		ipParts := s.ipManager.getIPv6Parts(ip)
		blackIPParts := s.ipManager.getIPv6Parts(blackListIP)

		for index, part := range ipParts {
			if blackIPParts[index] == "*" || blackIPParts[index] == part {
				match++
			}
		}
	}

	return match == 8
}