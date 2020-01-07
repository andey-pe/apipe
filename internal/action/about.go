package action

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func About(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "About page!\n")
}
