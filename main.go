package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	controller "./controllers"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/teste", controller.GetTeste)
	fmt.Println("Servidor está rodando na porta 8080")
	router.Run()
}