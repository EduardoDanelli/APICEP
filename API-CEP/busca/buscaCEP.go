package busca

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin" // gin
)

type Cidade struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func BuscaCEP(c *gin.Context) {

	cep := c.Query("cep")

	if len(cep) < 8 {
		c.JSON(http.StatusOK, gin.H{"error": "CEP inválido, faltam números"})
		return
	}

	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao conectar com o ViaCEP"})
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%s\n", data)

	var cidadeinfo Cidade
	if err := json.Unmarshal([]byte(data), &cidadeinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sb := string(body)
	log.Printf(sb)

	c.JSON(http.StatusOK, cidadeinfo)
}

func BuscaCEPABERTO(c *gin.Context) {

	resp, err := http.Get("https://cepaberto.com/ws/" + cep + "/json")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao conectar com CEPABERTO"})
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()
	fmt.Printf("%s\n", data)

	var cidadeinfo Cidade
	if err := json.Unmarshal([]byte(data), &cidadeinfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sb := string(body)
	log.Printf(sb)
}
