package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mammenj/go-health/models"
)

var proc_store models.ProcStore

func init() {
	proc_store = models.NewSqlliteProcStore()
}

func postMedProcs(c *gin.Context) {
	// Validate input

	var input models.MedProcs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medproc := models.MedProcs{Name: input.Name, Cost: input.Cost, Desription: input.Desription}
	ID, err := proc_store.Create(medproc)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "A new Practice Record " + ID + "Created!"})
}

func readMedProcs(c *gin.Context) {

	medprocs, err := proc_store.Get()
	checkErr(err)

	if medprocs == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	} else {
		//queues.Consume()
		c.JSON(200, gin.H{"practices": medprocs})
	}
}

func updateMedProcs(c *gin.Context) {
	// Validate input
	userid := c.Param("id")
	id, perr := strconv.Atoi(userid)
	if perr != nil {
		c.JSON(500, gin.H{"error": perr.Error()})
		return
	}

	var input models.MedProcs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	medproc := models.MedProcs{Name: input.Name, Cost: input.Cost, Desription: input.Desription, Id: id}
	err := proc_store.Update(medproc)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "Practices Record Updated!", "ID": id})

}

func deleteMedProcs(c *gin.Context) {

	if err := proc_store.Delete(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
