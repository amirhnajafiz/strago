package internal

import "time"

// service
// keeps the data of a service that
// has been given to strago.
type service struct {
	// service ip.
	ip string
	// number of requests that are sent to this service.
	used int
	// busy time of this service
	busy time.Duration
}
