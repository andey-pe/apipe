package route

import (
	"fmt"

	"github.com/andey-pe/apipe/internal/core/router"
	"github.com/valyala/fasthttp"
)

// ConfigureRoute ..
func ConfigureRoute(r router.RouterAbstract) {
	r.Get("/", home, AuthMiddleware, OAuthMiddleware)
	r.Get("/about", about)
}

func home(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func about(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "About page!\n")
}
