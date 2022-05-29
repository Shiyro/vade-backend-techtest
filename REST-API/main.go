package main

import (
	"errors"
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
	//On retourne toutes les donnees
	context.JSON(http.StatusOK, documents)
}

func getDocumentsById(id string) (*document, error) {
	//Retourne un document par ID
	for i, t := range documents {
		if t.ID == id {
			return &documents[i], nil
		}
	}
	return nil, errors.New("document not found")
}

func getDocument(context *gin.Context) {
	//Retourne le document recherche par ID
	id := context.Param("id")
	document, err := getDocumentsById(id)

	if err != nil {
		//Si il n'y a pas de documents correspondant a l'ID, on affiche qu'il est introuvable
		context.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
		return
	}

	context.JSON(http.StatusOK, document)
}

func addDocument(context *gin.Context) {
	var newDocument document

	//On essaye de convertir les donnees recu en JSON
	if err := context.BindJSON(&newDocument); err != nil {
		return
	}

	//Verifie si un document avec le meme id existe
	if document, err := getDocumentsById(newDocument.ID); document != nil && err == nil {
		context.JSON(http.StatusConflict, gin.H{"message": "Document with same id already exist"})
		return
	}

	documents = append(documents, newDocument)
	context.JSON(http.StatusCreated, newDocument)
}

func removeDocument(context *gin.Context) {
	//Supprime un document en fonction de l'ID

	id := context.Param("id")

	for i, t := range documents {
		if t.ID == id {
			//On retire le doc de la liste si il est trouve
			documents[i] = documents[len(documents)-1]
			documents = documents[:len(documents)-1]
			context.JSON(http.StatusOK, t)
			return
		}
	}

	//Si le doc n'existe pas on affiche qu'il est introuvable
	context.JSON(http.StatusNotFound, gin.H{"message": "Document not found"})
}

func main() {
	router := gin.Default()
	router.GET("/documents", getDocuments)
	router.GET("/documents/:id", getDocument)
	router.POST("/documents", addDocument)
	router.DELETE("/documents/:id", removeDocument)
	router.Run("localhost:8080")
}
