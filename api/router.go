package api

import (
	handler "Hackathon_Backend/api/handlers"
	_ "Hackathon_Backend/api/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())

	router.Use(cors.Default()) //allows all origin
	v1 := router.Group("/api/v1/")
	{

		users_api := v1.Group("/user/")
		{
			users_api.POST("/SignUp", handler.SignUp)
			users_api.POST("/login", handler.SignIn)
		}

		/*	item_api := v1.Group("/item/")
			{
				item_api.GET("/:id", middleware.TokenAuthMiddleware("customer:read"), handler.GetItemById)
				item_api.GET("/list-user/:id", middleware.TokenAuthMiddleware("customer:read"), handler.GetItemListByMenuId)
				item_api.POST("/create", middleware.TokenAuthMiddleware("admin:write"), handler.CreateItem)
				item_api.PUT("/update/:id", middleware.TokenAuthMiddleware("admin:write"), handler.UpdateItem)
				item_api.DELETE("/delete/:id", middleware.TokenAuthMiddleware("admin:write"), handler.DeleteItem)
			}*/
	}

	return router
}
