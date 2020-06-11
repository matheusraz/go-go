package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	controller "./controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/teste", controller.Teste)
	fmt.Println("Servidor está rodando na porta 8080")
	router.Run(":9090")
}