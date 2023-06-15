package main

import (
	server "computers_service/internal/grpc_server"
	"computers_service/internal/rabbit"
)

func main() {
	server.Start()
	rabbit.StartRabbitServer()
}
