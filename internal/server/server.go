package server

import (
	"fmt"
	"github.com/andey-pe/apipe/internal/config"
	"github.com/andey-pe/apipe/internal/router"
	"github.com/valyala/fasthttp"
	"log"
)

type ServerAbstract interface {
	Run()
}

type Server struct {
	config *config.Config
	router router.RouterAbstract
}

func NewServer(config *config.Config, router router.RouterAbstract) ServerAbstract {
	return &Server{
		config: config,
		router: router,
	}
}

func (s *Server) Run() {
	s.configRouter()
	log.Println("router configured")

	log.Printf("server starting on port %v", s.config.Server.Address)
	fasthttp.ListenAndServe(fmt.Sprintf(":%v", s.config.Server.Address), s.router.GetHandle())
}

func (s *Server) configRouter() {
	s.router.Get("/", home, AuthMiddleware)
	s.router.Get("/about", about)
}

func home(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func about(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "About page!\n")
}
