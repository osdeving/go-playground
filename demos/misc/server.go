//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ViaCEP struct {
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

// Como usar:
// 1. Execute o servidor: go run server.go
// 2. Acesse no navegador ou via curl: http://localhost:8080/cep/01001000
// Exemplo curl:
// curl http://localhost:8080/cep/01001000
func handleCEP(w http.ResponseWriter, r *http.Request) {

	// obtem tudo após o /cep/
	// ex: se chegar /cep/01001000 ele retorna 01001000
	cep := strings.TrimPrefix(r.URL.Path, "/cep/")

	// get param
	// cep := r.URL.Query().Get("cep")
		
	// faz a chamada para a API do ViaCEP usando o cep obtido
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	res := c(http.Get(url))
	defer res.Body.Close()

	
	// lê o corpo da resposta da API usando io.ReadAll
	body := c(io.ReadAll(res.Body))
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// decodifica o JSON da resposta da API usando json.Unmarshal
	var cepData ViaCEP
	void(json.Unmarshal(body, &cepData))

	// retorna o JSON da resposta da API setando o Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// converte para json já escrevendo no response writer
	json.NewEncoder(w).Encode(cepData)

	// // converte para json e retorna
	// jsonData, err := json.Marshal(cepData)
	// if err != nil {
	// 	panic(err)
	// }
	// w
	// w.Write(jsonData)
}

func main() {
	http.HandleFunc("/cep/", handleCEP)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

// must executa uma função que retorna um erro e faz panic caso ocorra erro
func c[T any](result T, err error) T {
	if err != nil {
		panic(err)
	}
	return result
}

func cr[T any](result T, err error, resolver func(error) T) T {
	if err != nil {
		if resolver != nil {
			return resolver(err)
		}
		panic(err)
	}
	return result
}

func void(err error) {
	if err != nil {
		panic(err)
	}
}

