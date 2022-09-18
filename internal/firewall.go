package internal

import (
	"strings"
)

// checkIPRangeInBlackList
// goes through blacklist ips and checks
// if the given ip is in the list or not.
func (s *server) checkIPRangeInBlackList(ip string) bool {
	match := 0

	for _, blackListIP := range s.blacklist {
		ipParts := strings.Split(ip, ".")
		blackIPParts := strings.Split(blackListIP, ".")

		for index, part := range ipParts {
			if blackIPParts[index] == "*" || blackIPParts[index] == part {
				match++
			}
		}
	}

	return match == 4
}
