package routes

import (
	"github.com/Vause/Coffee-Rating-api/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()
	r.Use(cors.Default())
	v1 := r.Group("/v1")
	{
		v1.GET("ratings", controllers.GetRatings)
		v1.GET("ratings/:id", controllers.GetRatingById)
		v1.POST("ratings", controllers.CreateRating)
		// v1.PUT("ratings/:id", controllers.UpdateRatingById)
		// v1.DELETE("ratings/:id", controllers.DeleteRatingById)
	}

	return r
}
