package user

import "github.com/gin-gonic/gin"

type ProductRouter struct{}

func (p *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/get_all_product")
		productRouterPublic.GET("/get_product_by_id")
	}

	// private router
	productRouterPrivate := Router.Group("/product")
	{
		productRouterPrivate.POST("/create_product")
		productRouterPrivate.POST("/update_product")
		productRouterPrivate.DELETE("/delete_product")
	}
}
