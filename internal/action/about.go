package action

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (a *Action) About(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "About page!\n")
}
