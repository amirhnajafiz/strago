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
