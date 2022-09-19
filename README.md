<p align="center">
  <img src="assets/strago.webp" alt="logo" />
</p>

<h1 align="center">
Strago
</h1>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    <img src="https://img.shields.io/badge/Version-0.0.3-green?style=for-the-badge&logo=none" alt="version" />
</p>

Simple traffic controller agent with Golang. With **Strago** you can create and config a load balancer
for your services. Load balancer logic of **Strago** is very simple, it works by the number of requests per service.
All you need to do is to give your service addresses as an IP and leave the rest up to **Strago**.

## How to use?
Install library:
```shell
go get -u github.com/amirhnajafiz/strago
```

### Example
If you set two _echo servers_ on localhost ports ```5050 and 5051```, then
you have to set the strago server like the example below:
```go
package main

import "github.com/amirhnajafiz/strago"

func main() {
	server := strago.NewServer(
		strago.WithServices(
			strago.WithDefaultConfigs(),
			"127.0.0.1:5050",
			"127.0.0.1:5051",
		),
	)

	if err := server.Start(); err != nil {
		panic(err)
	}
}
```

### Firewall
You can ban any ip range (IPv4 or IPv6):
```go
...
// ban any ip in range 127... version 4
if err := server.BanIP("127.*.*.*"); err != nil {
    panic(err)
}

// ban any ip in range 12...12 version 6
if err := server.BanIP("12:*:*:*:22:*:*:*"); err != nil {
    panic(err)
}

// allow requests passing
server.Enable()
...
```

### Test
You can test the above code by creating two _echo servers_:
```shell
### generating a service on port 5050
go run example/echo/main.go 5050
### generating a service on port 5051
go run example/echo/main.go 5051
```

Now you can test the load-balancer:
```shell
curl localhost:9370
```
