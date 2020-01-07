package main

import (
	"github.com/andey-pe/apipe/internal/config"
	"github.com/andey-pe/apipe/internal/router"
	"github.com/andey-pe/apipe/internal/server"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	c := dig.New()

	c.Provide(config.NewConfig)
	c.Provide(router.NewRouter)
	c.Provide(server.NewServer)

	return c
}
