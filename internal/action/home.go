package action

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Home(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}
