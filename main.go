package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mammenj/go-health/handlers"
	"github.com/mammenj/go-health/models"
)

func main() {
	r := gin.Default()

	proc_handler := handlers.CreateProcslHandler(models.NewSqlliteProcStore())
	proc_router := r.Group("/procs")
	{
		proc_router.POST("/", proc_handler.PostMedProcs)
		proc_router.PUT("/:id", proc_handler.UpdateMedProcs)
		proc_router.GET("/", proc_handler.ReadMedProcs)
		proc_router.DELETE("/:id", proc_handler.DeleteMedProcs)
	}
	hospital_handler := handlers.CreateHospitalHandler(models.NewSqlliteHospitalStore())

	hospital_router := r.Group("/hospital")
	{
		hospital_router.POST("/", hospital_handler.CreateHospital)
		hospital_router.PUT("/:id", hospital_handler.UpdateHospital)
		hospital_router.GET("/", hospital_handler.GetHospitals)
		hospital_router.DELETE("/:id", hospital_handler.DeleteHospital)
	}

	// By default it serves on :8080
	r.Run()

}
