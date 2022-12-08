package api

import (
	handler "Hackathon_Backend/api/handlers"
	"Hackathon_Backend/api/middleware"
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
			users_api.POST("/basicInfo", middleware.TokenAuthMiddleware(), handler.AddBasicInfo)
			users_api.GET("/basicInfo", middleware.TokenAuthMiddleware(), handler.GetBasicInfo)
			users_api.PUT("/basicInfo", middleware.TokenAuthMiddleware(), handler.UpdateBasicInfo)
			users_api.POST("/measurement", middleware.TokenAuthMiddleware(), handler.AddMeasurement)
			users_api.GET("/measurement/:type", middleware.TokenAuthMiddleware(), handler.GetMeasurementByType)
			users_api.GET("/allMeasurement/:date", middleware.TokenAuthMiddleware(), handler.GetAllMeasurementByDate)
			users_api.GET("/allMeasurementOfDate/:date", middleware.TokenAuthMiddleware(), handler.GetAllMeasurementOfDate)
			users_api.POST("/upload", middleware.TokenAuthMiddleware(), handler.UploadFile)
			users_api.GET("/download", middleware.TokenAuthMiddleware(), handler.DownloadFile)
			users_api.POST("/visit", middleware.TokenAuthMiddleware(), handler.AddVisit)
			users_api.GET("/allVisit", middleware.TokenAuthMiddleware(), handler.GetAllVisit)
		}

	}

	return router
}
