package models

type Cidade struct {
	Cep         string       `json:"cep"`
	Logradouro  string       `json:"logradouro"`
	Complemento string       `json:"complemento"`
	Bairro      string       `json:"bairro"`
	Localidade  string       `json:"localidade"`
	Uf          string       `json:"uf"`
	Ibge        string       `json:"ibge"`
	Gia         string       `json:"gia"`
	Ddd         string       `json:"ddd"`
	Siafi       string       `json:"siafi"`
	Erro        bool         `json:"erro"`
	Cidade      CidadeAberto `json:"cidade"`
}

type CidadeAberto struct {
	Cidade    CidadeCidadeAberto `json:"cidade"`
	Estado    Estado             `json:"estado"`
	Altitude  float64            `json:"altitude"`
	Cep       string             `json:"cep"`
	Latitude  string             `json:"latitude"`
	Longitude string             `json:"longitude"`
}
type CidadeCidadeAberto struct {
	Ddd  int    `json:"ddd"`
	Nome string `json:"nome"`
}
type Estado struct {
	Sigla string `json:"sigla"`
}
