package main

import (
	busca "cep/busca/cep"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/cep", busca.BuscaCEP)

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
