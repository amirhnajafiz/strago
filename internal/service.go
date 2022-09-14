package internal

type service struct {
	enable bool
	ip     string
	used   int
}

func createServices(services []string) []*service {
	list := make([]*service, len(services))

	for _, s := range services {
		list = append(list, &service{
			enable: true,
			ip:     s,
			used:   0,
		})
	}

	return list
}
