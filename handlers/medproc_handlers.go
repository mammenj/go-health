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

type ProcslHandler struct {
	store models.ProcStore
}

func NewSqlliteProcStore() *models.ProcSqlliteStore {
	envs := util.LoadEnv()
	log.Println("ENVS: ", envs)

	db, err := sql.Open(envs["DB_DIALECT_PROCS"], envs["DB_PROCS"])
	if err != nil {
		panic(err)
	}
	return &models.ProcSqlliteStore{
		DB: db,
	}
}

func CreateProcslHandler(store models.ProcStore) *ProcslHandler {
	return &ProcslHandler{store: store}
}

func (p ProcslHandler) PostMedProcs(c *gin.Context) {
	// Valiate input

	var input models.MedProcs
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	medproc := models.MedProcs{Name: input.Name, Cost: input.Cost, Desription: input.Desription}
	ID, err := p.store.Create(medproc)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "A new Practice Record " + ID + "Created!"})
}

func (p ProcslHandler) ReadMedProcs(c *gin.Context) {

	medprocs, err := p.store.Get()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	if medprocs == nil {
		c.JSON(404, gin.H{"error": "No Records Found"})
		return
	}
	//queues.Consume()
	c.JSON(200, gin.H{"practices": medprocs})
}

func (p ProcslHandler) UpdateMedProcs(c *gin.Context) {
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
	err := p.store.Update(medproc)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return

	}
	c.JSON(200, gin.H{"message": "Practices Record Updated!", "ID": id})

}

func (p ProcslHandler) DeleteMedProcs(c *gin.Context) {

	if err := p.store.Delete(c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Deleted ID: " + c.Param("id")})
}
