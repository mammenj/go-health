package main

import (
	"github.com/gin-gonic/gin"

	"github.com/mammenj/go-health/handlers"
)

func main() {
	r := gin.Default()
	// API v1
	proc_router := r.Group("/procs")
	{
		proc_router.POST("/", handlers.PostMedProcs)
		proc_router.PUT("/:id", handlers.UpdateMedProcs)
		proc_router.GET("/", handlers.ReadMedProcs)
		proc_router.DELETE("/:id", handlers.DeleteMedProcs)
	}

	hospital_router := r.Group("/hospital")
	{
		hospital_router.POST("/", handlers.CreateHospital)
		hospital_router.PUT("/:id", handlers.UpdateHospital)
		hospital_router.GET("/", handlers.GetHospitals)
		hospital_router.DELETE("/:id", handlers.DeleteHospital)
	}

	// By default it serves on :8080
	r.Run()

}
