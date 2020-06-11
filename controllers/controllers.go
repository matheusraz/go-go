package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"../model"
)
	

func Teste(response *gin.Context) {
	usuario := model.User{}
	response.BindJSON(&usuario)
	response.JSON(http.StatusOK, gin.H{"status": 1, "msg": usuario})
}