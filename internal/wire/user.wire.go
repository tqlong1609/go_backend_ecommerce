//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
	"github.com/tqlong1609/go_backend_ecommerce/internal/repo"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
)

func InitUserRouterHandler() (*controllers.UserController, error) {
	wire.Build(
		repo.InitUserRepository,
		services.InitUserService,
		controllers.InitUserController,
	)
	return new(controllers.UserController), nil
}
