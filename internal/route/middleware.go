package route

import (
	"log"

	"github.com/valyala/fasthttp"
)

// AuthMiddleware ..
func AuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Println("auth middlewares")

		h(ctx)
	}
}

// OAuthMiddleware ..
func OAuthMiddleware(h fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		log.Println("OAuthMiddleware middlewares")

		h(ctx)
	}
}
