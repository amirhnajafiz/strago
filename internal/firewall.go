package internal

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

const (
	ipV6          = 6
	ipV4          = 4
	unsupportedIp = 0
)

// ipType
// manages to find the ip version of a given ip.
func (s *server) ipType(ip string) int {
	if govalidator.IsIPv6(ip) {
		return ipV6
	} else if govalidator.IsIPv4(ip) {
		return ipV4
	} else {
		return unsupportedIp
	}
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
