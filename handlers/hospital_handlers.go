package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mammenj/go-health/models"
	"github.com/mammenj/go-health/util"
)

type HospitalHandler struct {
	store models.HospitalStore
}

func NewSqlliteHospitalStore() *models.HospitalSqlliteStore {
	envs := util.LoadEnv()
	log.Println("ENVS: ", envs)

	db, err := sql.Open(envs["DB_DIALECT_HOSPITAL"], envs["DB_HOSPITAL"])
	if err != nil {
		panic(err)
	}
	return &models.HospitalSqlliteStore{
		DB: db,
	}
}

func CreateHospitalHandler(store models.HospitalStore) *HospitalHandler {
	return &HospitalHandler{store: store}
}

func (h *HospitalHandler) CreateHospital(c *gin.Context) {
	// Validate input
	var input models.Hospital
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hospital := models.Hospital{Name: input.Name, Address: input.Address, City: input.City, Country: input.Country, Pin: input.Pin}
	ID, err := h.store.Create(hospital)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "A new Hosipital ID: " + ID + "Created!"})
}

func (h *HospitalHandler) GetHospitals(c *gin.Context) {

	hospitals, err := h.store.Get()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}

	if hospitals == nil {
		c.JSON(404, gin.H{"error": "No Hospital Records Found"})
		return
	}
	//queues.Consume()
	c.JSON(200, gin.H{"hospitals": hospitals})

}

func (h *HospitalHandler) UpdateHospital(c *gin.Context) {
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
	err := h.store.Update(hospital)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "Hospital Record Updated!", "ID: ": id})

}

func (h *HospitalHandler) DeleteHospital(c *gin.Context) {
	// Get model if exist
	if err := h.store.Delete(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Hospital ID: " + c.Param("id") + " Record Deleted!"})
}
