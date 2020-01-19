package action

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func (a *Action) Home(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}
