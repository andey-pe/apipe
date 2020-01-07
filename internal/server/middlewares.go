package server

import "github.com/valyala/fasthttp"

import "log"

func AuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Println("auth middlewares")

		h(ctx)
	}
}
