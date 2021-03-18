package busca

import (
	"cep/busca/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BuscaCEPaberto(c *gin.Context) {

	cep := c.Query("cepAberto")

	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, ".", "")

	// Verificando se o CEP é valido.
	if len(cep) != 8 {
		c.JSON(400, gin.H{"error": "CEP inválido"})
		return
	}

	url := "https://www.cepaberto.com/api/v3/cep?cep=" + cep
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token token=adee59a89420858065b72d17ce8c42b6")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Erro: %v", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Erro ao ler o corpo da requisição: %v", err)
	}

	var cidadeinfo models.CidadeAberto
	if err := json.Unmarshal([]byte(body), &cidadeinfo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Erro ao conectar ao CepAberto"})
		return
	}
	c.JSON(http.StatusOK, cidadeinfo)
	fmt.Print(cep)

}
