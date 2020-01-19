package route

import (
	"github.com/andey-pe/apipe/internal/action"
	"github.com/andey-pe/apipe/internal/core/router"
)

// ConfigureRoute ..
func ConfigureRoute(r router.RouterAbstract, a *action.Action) {
	r.Get("/", a.Home, AuthMiddleware, OAuthMiddleware)
	r.Get("/about", a.About)
}
