package main

import (
	"flag"
	"log"

	"github.com/ebcardoso/api-rest-golang/server"
)

func main() {
	//Listen Port
	port := flag.String("port", ":8080", "the server address")
	flag.Parse()

	//Creating new server
	server, err := server.NewServer(*port)
	if err != nil {
		log.Fatal(err)
		return
	}

	//Starting Server
	log.Printf("Server running on port: %s", *port)
	log.Fatal(server.StartServer())
}
