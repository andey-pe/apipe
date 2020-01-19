package main

import (
	"github.com/andey-pe/apipe/internal/core/config"
	"github.com/andey-pe/apipe/internal/action"
	"github.com/andey-pe/apipe/internal/core/router"
	"github.com/andey-pe/apipe/internal/core/server"
	"go.uber.org/dig"
)

func buildContainer() *dig.Container {
	c := dig.New()

	c.Provide(config.NewConfig)
	c.Provide(router.NewRouter)
	c.Provide(action.NewAction)
	c.Provide(server.NewServer)

	return c
}
