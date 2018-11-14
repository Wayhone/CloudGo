package main

import (
	"CloudGo/service"
	flag "github.com/spf13/pflag"
	"os"
)

const (
	PORT string = "8080" // default port 8080
)

func main() {
	port := os.Getenv("PORT") // check the port
	if len(port) == 0 {
		port = PORT // use the default port if no port specified
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for http listening")	// check the argument
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	service.NewServer(port)
}