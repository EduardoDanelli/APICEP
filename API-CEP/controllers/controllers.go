package controllers

import "github.com/gin-gonic/gin"

func CEP(c *gin.Context) {
	// Verificando se o CEP é valido.
	if len(cep) != 8 {
		c.AbortWithStatusJSON(400, gin.H{"error": "CEP inválido"})
		return
	}

	// Tentando buscar o CEP no CepAberto
	cepInfo, _ := services.GetCepAberto(cep)
	if cepInfo != nil {
		c.JSON(200, gin.H{"data": cepInfo})
		return
	}

	// Tentando buscar o CEP no ViaCep
	cepInfo, _ = services.GetViaCep(cep)
	if cepInfo != nil {
		c.JSON(200, gin.H{"data": cepInfo})
		return
	}

	c.AbortWithStatusJSON(400, gin.H{"error": "CEP não encontrado"})
}
