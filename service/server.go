package service

import (
	"github.com/go-martini/martini"
)

func NewServer(port string) {
	// get an instance of martini
	m := martini.Classic()

	// receive the request of a GET method
	m.Get("/", func(params martini.Params) string {
		return "Welcome to Cloud_Go!"
	})

	// run the server
	m.RunOnAddr(":" + port)
}