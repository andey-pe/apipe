package server

import (
	"fmt"

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
}

// NewServer __constructor
func NewServer(config *config.Config, router router.RouterAbstract) ServerAbstract {
	return &Server{
		config: config,
		router: router,
	}
}

// Run server
func (s *Server) Run() {
	route.ConfigureRoute(s.router)
	fasthttp.ListenAndServe(fmt.Sprintf(":%v", s.config.Server.Address), s.router.GetHandle())
}
