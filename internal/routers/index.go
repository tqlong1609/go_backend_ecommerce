package routers

import (
	"github.com/tqlong1609/go_backend_ecommerce/internal/routers/admin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/routers/user"
)

type RouterGroup struct {
	User  user.UserRouterGroup
	Admin admin.AdminRouterGroup
}

var RouterGroupApp = new(RouterGroup)
