package router

import (
	"esave/database/dao"
	services "esave/services/dtos"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func PingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}

func CreateSpentHandler(c *gin.Context) {
	var spentDto services.SpentDto
	if err := c.ShouldBindJSON(&spentDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	dao.SpentDaoImpl{}.CreateSpent(spentDto)
	c.JSON(http.StatusOK, spentDto)
	return
}

func LandingHandler(c *gin.Context) {
	c.String(http.StatusOK, "Wellcome to esave")
}

func GetSpentByIdHandler(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 64)
	c.JSON(http.StatusOK, dao.SpentDaoImpl{}.GetSpentById(uintId))
	return
}

func DeleteSpentByIdHandler(c *gin.Context) {
	id := c.Param("id")
	uintId, _ := strconv.ParseUint(id, 10, 64)
	dao.SpentDaoImpl{}.DeleteSpentById(uintId)
}

func UpdateSpentHandler(c *gin.Context) {
	var spentDto services.SpentDto
	if err := c.ShouldBindJSON(&spentDto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	dao.SpentDaoImpl{}.UpdateSpent(spentDto)
	c.JSON(http.StatusOK, spentDto)
	return
}

func SetupRouter() *gin.Engine {
	// Crear una instancia de Gin
	log.Println("Starting router...")
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,                                                // Dominios permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Métodos permitidos
		AllowHeaders:     []string{"Content-Type"},                            // Cabeceras permitidas
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Rutas y manejadores
	router.GET("/ping", PingHandler)
	router.GET("/", LandingHandler)
	router.POST("/spent/create", CreateSpentHandler)
	router.GET("/spent/:id", GetSpentByIdHandler)
	router.DELETE("/spent/:id", DeleteSpentByIdHandler)
	router.PUT("/spent/update", UpdateSpentHandler)

	return router
}
