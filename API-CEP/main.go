package main

import (
	"cep/busca"
	//busca "cep/busca/cep"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/cep", busca.BuscaCEP)
	r.GET("/cepAberto", busca.BuscaCEPaberto)

	if err := r.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
