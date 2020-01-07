package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type RouterAbstract interface {
	Get(string, fasthttp.RequestHandler, ...func(fasthttp.RequestHandler) fasthttp.RequestHandler)
	Post(string, fasthttp.RequestHandler)
	Delete(string, fasthttp.RequestHandler)
	GetHandle() fasthttp.RequestHandler
}

type Router struct {
	r *fasthttprouter.Router
}

func NewRouter() RouterAbstract {
	return &Router{
		r: fasthttprouter.New(),
	}
}

// Get method
func (r *Router) Get(path string, handle fasthttp.RequestHandler, middlewares ...func(fasthttp.RequestHandler) fasthttp.RequestHandler) {
	r.r.GET(path, withMiddlewares(handle, middlewares))
}

// Post method
func (r *Router) Post(path string, handle fasthttp.RequestHandler) {
	r.r.POST(path, handle)
}

// Delete method
func (r *Router) Delete(path string, handle fasthttp.RequestHandler) {
	r.r.DELETE(path, handle)
}

// GetHandle return router handler
func (r *Router) GetHandle() fasthttp.RequestHandler {
	return r.r.Handler
}

func withMiddlewares(handle fasthttp.RequestHandler, middlewares []func(fasthttp.RequestHandler) fasthttp.RequestHandler) fasthttp.RequestHandler {
	for _, m := range middlewares {
		handle = m(handle)
	}

	return handle
}
