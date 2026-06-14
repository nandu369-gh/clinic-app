package routes

import (
	"clinic-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default() // Includes Logger and Recovery middleware
	// API Group V1
	v1 := r.Group("/api/v1")
	{
		v1.POST("/appointments", controllers.CreateAppointment)
		// Future endpoints to build:
		// v1.GET("/appointments",controllers.GetAllAppointments)
		// v1.GET("/appointments/:id",controllers.GetAppointmentByID)
	}
	return r
}
