package internal

import (
	"strings"

	"github.com/asaskevich/govalidator"
)

const (
	ipV6          = 6
	ipV4          = 4
	unsupportedIp = 0

	ipV6Separator = ":"
	ipV4Separator = "."
)

// ipManager
// stores all data for firewall ips.
type ipManager struct {
	ipV4List []string
	ipV6List []string
}

// ipType
// manages to find the ip version of a given ip.
func (ipm *ipManager) ipType(ip string) int {
	if govalidator.IsIPv6(ip) {
		return ipV6
	} else if govalidator.IsIPv4(ip) {
		return ipV4
	} else {
		return unsupportedIp
	}
}

// genericIPType
// checks for input ip type.
// this method is used for generic ips that are set
// for firewall black lists.
func (ipm *ipManager) genericIPType(ip string) int {
	var parts []string

	parts = strings.Split(ip, ipV4Separator)
	if len(parts) == 4 {
		return ipV4
	}

	parts = strings.Split(ip, ipV6Separator)
	if len(parts) == 6 {
		return ipV6
	}

	return unsupportedIp
}

// getIPv4Parts
// returns all parts of an IP version 4.
func (ipm *ipManager) getIPv4Parts(ip string) []string {
	return strings.Split(ip, ipV4Separator)
}

// getIPv6Parts
// returns all parts of an IP version 6.
func (ipm *ipManager) getIPv6Parts(ip string) []string {
	return strings.Split(ip, ipV6Separator)
}

// validateIPv4
// custom validation for IP version 4.
func (ipm *ipManager) validateIPv4(ip string) bool {
	success := 0

	for _, part := range ipm.getIPv6Parts(ip) {
		if part == "*" || govalidator.IsNumeric(part) {
			success++
		}
	}

	return success == 4
}

// validateIPv6
// custom validation for IP version 6.
func (ipm *ipManager) validateIPv6(ip string) bool {
	success := 0

	for _, part := range ipm.getIPv6Parts(ip) {
		if part == "*" || govalidator.IsHexadecimal(ip) {
			success++
		}
	}

	return success == 8
}

// addToBlacklist
// add on ip to blacklist based on its version.
func (ipm *ipManager) addToBlacklist(ip string) bool {
	switch ipm.genericIPType(ip) {
	case ipV6:
		if !ipm.validateIPv6(ip) {
			return false
		}

		ipm.ipV6List = append(ipm.ipV6List, ip)

		return true
	case ipV4:
		if !ipm.validateIPv4(ip) {
			return false
		}

		ipm.ipV4List = append(ipm.ipV4List, ip)

		return true
	}

	return false
}

// removeFromBlacklist
// removing a given ip from blacklist.
func (ipm *ipManager) removeFromBlacklist(ip string) bool {
	var result bool

	switch ipm.genericIPType(ip) {
	case ipV6:
		ipm.ipV6List, result = removeFromList(ipm.ipV6List, ip)
	case ipV4:
		ipm.ipV4List, result = removeFromList(ipm.ipV4List, ip)
	}

	return result
}

// removeFromList
// gets a list and removes one given item from it, if it exists.
func removeFromList(list []string, ip string) ([]string, bool) {
	for index, item := range list {
		if item == ip {
			list = append(list, list[index+1:]...)

			return list, true
		}
	}

	return list, false
}
