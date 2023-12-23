package main

import (
	"kresimir.dev/modules/server"
	"log"
)

const PORT = 8080

func main() {
	log.Println("Starting on port: ", PORT)
	server := server.CreateServer(PORT)
	server.Listen()
}
