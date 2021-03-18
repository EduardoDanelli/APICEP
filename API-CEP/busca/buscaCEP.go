package busca

import (
	"cep/busca/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin" // gin
)

func BuscaCEP(c *gin.Context) {

	cep := c.Query("cep")
	cep = strings.ReplaceAll(cep, "-", "")
	cep = strings.ReplaceAll(cep, ".", "")

	// Verificando se o CEP é valido.
	if len(cep) != 8 {
		c.JSON(400, gin.H{"error": "CEP inválido"})
		return
	}

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "Erro ao conectar com o ViaCEP"})
		return
	}

	if resp.StatusCode != 200 {
		c.JSON(http.StatusOK, gin.H{"error": "Erro ao conectar com ViaCEP"})
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler resposta"})
		return
	}

	defer resp.Body.Close()
	fmt.Printf("%s\n", data)

	var cidadeinfo models.Cidade
	if err := json.Unmarshal([]byte(data), &cidadeinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao conectar com ViaCEP"})
		return
	}

	/*if cidadeinfo.Erro {
		c.JSON(http.StatusOK, gin.H{"error": "CEP não encontrado"})
		return
	}*/

	c.JSON(http.StatusOK, cidadeinfo)
	fmt.Print(cep)
}
