package internal

import "log"

type service struct {
	enable bool
	ip     string
	used   int
}

func createServices(services []string) []*service {
	list := make([]*service, len(services))

	for index, s := range services {
		list[index] = &service{
			enable: true,
			ip:     s,
			used:   0,
		}
	}

	log.Println(len(list))
	log.Println(len(services))

	return list
}
