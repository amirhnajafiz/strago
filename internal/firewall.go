package internal

import (
	"fmt"
	"strings"
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
