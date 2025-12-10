package main

import (
	"xi/cmd/server/cmd"
	"xi/internal/service"
)

func main() {
	service.InitPre() // Init services

	cmd.Init()
}
