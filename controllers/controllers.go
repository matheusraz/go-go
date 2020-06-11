package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
	

func GetTeste(response *gin.Context) {
	response.JSON(http.StatusOK, gin.H{"status": 1, "msg": "Funfou Caraiou!!"})
}