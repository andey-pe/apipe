package main

import "github.com/andey-pe/apipe/internal/server"

func main() {
	c := buildContainer()

	c.Invoke(func(s server.ServerAbstract) {
		s.Run()
	})
}
