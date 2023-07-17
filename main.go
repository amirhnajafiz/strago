package main

import "flag"

func main() {
	var (
		Services    = flag.String("services", "", "list of services")
		Port        = flag.Int("port", 8080, "http port")
		Secure      = flag.Bool("secure", false, "http secure or not")
		BalanceType = flag.Int("type", 1, "load balancing type")
	)

	flag.Parse()
}
