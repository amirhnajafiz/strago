# from golang alpine image
FROM golang:1.20-alpine as builder

# app work directory
WORKDIR /app

# copy go.mod and go.sum
COPY go.mod go.sum ./

# download deps
RUN go mod download

# copy all files
COPY . .

# building go executable file
RUN go build -o main

# second stage
FROM alpine

# create src work directory
WORKDIR /src

# copy execute file
COPY --from=builder /app/main main

# start strago service
CMD ./main