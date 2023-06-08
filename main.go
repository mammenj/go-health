package main

import (
	"log"

	"github.com/gin-gonic/gin"

	//"github.com/mammenj/go-health/queues"
	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := gin.Default()
	// API v1
	proc_router := r.Group("/procs")
	{
		proc_router.POST("/", postMedProcs)
		proc_router.PUT("/:id", updateMedProcs)
		proc_router.GET("/", readMedProcs)
		proc_router.DELETE("/:id", deleteMedProcs)
	}

	hospital_router := r.Group("/hospital")
	{
		hospital_router.POST("/", createHospital)
		hospital_router.PUT("/:id", updateHospital)
		hospital_router.GET("/", getHospitals)
		hospital_router.DELETE("/:id", deleteHospital)
	}

	// By default it serves on :8080
	r.Run()

}
