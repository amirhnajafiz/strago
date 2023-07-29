<p align="center">
  <img src="assets/strago.webp" alt="logo" />
</p>

<h1 align="center">
    Strago
</h1>

<p align="center">
    <img src="https://img.shields.io/badge/Go-1.20-00ADD8?style=for-the-badge&logo=go" alt="go version" />
    <img src="https://img.shields.io/badge/Version-0.1.1-green?style=for-the-badge&logo=github" alt="version" />
</p>

Simple traffic controller agent with Golang. With **Strago** you can create and config a load balancer
for your services. Load balancing logic of **Strago** is very simple, it works by the number of requests per service
or the total burst time of requests in each service. 
All you need to do is to give your service addresses as an IP and leave the rest up to **Strago**.

## How to use?

Install library:

```shell
go get -u github.com/amirhnajafiz/strago
```

### build

Then build the executable file:

```shell
go build -o strago
chmod +x ./strago
STRAGO_PORT=9370 STRAGO_SERVICES='127.0.0.1:5050&127.0.0.1:5051' STRAGO_DEBUG=true ./main
```

### docker

Start strago on docker with:

```shell
docker pull amirhossein21/strago:v0.1.1
docker run -d \
  -e STRAGO_PORT=9370 \
  -e STRAGO_SERVICES='127.0.0.1:5050&127.0.0.1:5051' \
  amirhossein21/strago:v0.1.1
```

## Example

If you set two _echo servers_ on localhost ports ```5050 and 5051```, then
you have to set the strago server like the example below:

### test

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

## Envs

- ```STRAGO_SERVICES```: seperated with ```&```. Example: ```127.0.0.1:5050&127.0.0.1:5051```
- ```STRAGO_PORT```: http port
- ```STRAGO_SECURE```: using https or not (ture/false)
- ```STRAGO_TYPE```: load balancing type (1 is request count / 2 is burst time)
- ```STRAGO_DEBUG```: used for debug mode (true/false)

## Metrics & Health

Strago metrics will be exposed on ```localhost:9370/metrics```. And you can check the system
health on ```localhost:9370/health```.
