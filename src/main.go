package main

import (
	"kresimir.dev/modules/core"
)

func main() {
	server := core.CreateServer(8080)
	server.Listen()
}
