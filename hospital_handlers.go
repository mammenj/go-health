package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mammenj/go-health/models"
)

var store models.HospitalStore

func init() {
	store = models.NewSqlliteStore()
}

func createHospital(c *gin.Context) {
	// Validate input
	var input models.Hospital
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital := models.Hospital{Name: input.Name, Address: input.Address, City: input.City, Country: input.Country, Pin: input.Pin}
	ID, err := store.Create(hospital)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "A new Hosipital Record " + ID + "Created!"})
}

func getHospitals(c *gin.Context) {

	hospitals, err := store.Get()
	checkErr(err)

	if hospitals == nil {
		c.JSON(404, gin.H{"error": "No Hospital Records Found"})
		return
	} else {
		//queues.Consume()
		c.JSON(200, gin.H{"hospitals": hospitals})
	}
}

func updateHospital(c *gin.Context) {
	// Validate input
	userid := c.Param("id")
	id, perr := strconv.Atoi(userid)
	if perr != nil {
		c.JSON(500, gin.H{"error": perr.Error()})
		return
	}

	var input models.Hospital
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hospital := models.Hospital{Name: input.Name, Address: input.Address, City: input.City, Country: input.Country, Pin: input.Pin, Id: id}
	err := store.Update(hospital)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "Hospital Record Updated!", "ID": id})

}

func deleteHospital(c *gin.Context) {
	// Get model if exist
	if err := store.Delete(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Hospital " + c.Param("id") + " Record Deleted!"})
}
