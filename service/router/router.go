package router

import (
	"employee/handlers/employee"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	// Create router
	router := gin.Default()

	// Create employee handler
	eh := employee.NewEmployeeHandler()

	// Register handlers
	router.POST("/employee", eh.CreateEmployee)
	router.GET("/employee", eh.GetEmployee)
	router.PUT("/employee", eh.UpdateEmployee)
	router.DELETE("/employee", eh.DeleteEmployee)
	return router
}
