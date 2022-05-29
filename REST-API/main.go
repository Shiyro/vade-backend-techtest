package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type document struct {
	ID          string `json:"id"`
	Nom         string `json:"nom"`
	Description string `json:"description"`
}

var documents = []document{
	//On prerempli au démarrage pour avoir quelques donnees
	{ID: "1", Nom: "Williams", Description: "Développeur Backend"},
	{ID: "2", Nom: "Martin", Description: "Développeur Frontend"},
	{ID: "3", Nom: "Harris", Description: "Administrateur système"},
}

func getDocuments(context *gin.Context) {
	//On retoune toutes les donnees
	context.JSON(http.StatusOK, documents)
}

func main() {
	router := gin.Default()
	router.GET("/documents", getDocuments)
	router.Run("localhost:8080")
}
