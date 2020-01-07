package route

import (
	"github.com/andey-pe/apipe/internal/action"
	"github.com/andey-pe/apipe/internal/core/router"
)

// ConfigureRoute ..
func ConfigureRoute(r router.RouterAbstract) {
	r.Get("/", action.Home, AuthMiddleware, OAuthMiddleware)
	r.Get("/about", action.About)
}
