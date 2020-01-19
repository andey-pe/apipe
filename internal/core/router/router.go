package router

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type MiddlewareType = func(fasthttp.RequestHandler) fasthttp.RequestHandler

// RouterAbstract ..
type RouterAbstract interface {
	Get(string, fasthttp.RequestHandler, ...MiddlewareType)
	Post(string, fasthttp.RequestHandler, ...MiddlewareType)
	Delete(string, fasthttp.RequestHandler, ...MiddlewareType)
	GetHandle() fasthttp.RequestHandler
}

// Router ..
type Router struct {
	r *fasthttprouter.Router
}

// NewRouter __constructor
func NewRouter() RouterAbstract {
	return &Router{
		r: fasthttprouter.New(),
	}
}

// Get method
func (r *Router) Get(path string, handle fasthttp.RequestHandler, middlewares ...MiddlewareType) {
	r.r.GET(path, withMiddlewares(handle, middlewares))
}

// Post method
func (r *Router) Post(path string, handle fasthttp.RequestHandler, middlewares ...MiddlewareType) {
	r.r.POST(path, withMiddlewares(handle, middlewares))
}

// Delete method
func (r *Router) Delete(path string, handle fasthttp.RequestHandler, middlewares ...MiddlewareType) {
	r.r.DELETE(path, withMiddlewares(handle, middlewares))
}

// GetHandle return router handler
func (r *Router) GetHandle() fasthttp.RequestHandler {
	return r.r.Handler
}

// Handle middlewares
func withMiddlewares(handle fasthttp.RequestHandler, middlewares []MiddlewareType) fasthttp.RequestHandler {
	for _, m := range middlewares {
		handle = m(handle)
	}

	return handle
}
