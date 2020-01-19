package server

import (
	"fmt"

	"github.com/andey-pe/apipe/internal/action"
	"github.com/andey-pe/apipe/internal/core/config"
	"github.com/andey-pe/apipe/internal/core/router"
	"github.com/andey-pe/apipe/internal/route"
	"github.com/valyala/fasthttp"
)

// ServerAbstract ..
type ServerAbstract interface {
	Run()
}

//Server ..
type Server struct {
	config *config.Config
	router router.RouterAbstract
	action *action.Action
}

// NewServer __constructor
func NewServer(config *config.Config, router router.RouterAbstract, action *action.Action) ServerAbstract {
	return &Server{
		config: config,
		router: router,
		action: action,
	}
}

// Run server
func (s *Server) Run() {
	route.ConfigureRoute(s.router, s.action)
	fasthttp.ListenAndServe(fmt.Sprintf(":%v", s.config.Server.Address), s.router.GetHandle())
}
