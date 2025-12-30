package main

import (
	"xi/cmd/web/cmd"
	"xi/internal/services"
)

func main() {
	services.InitPre() // Init services

	cmd.Init()
}
