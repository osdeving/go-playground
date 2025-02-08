//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/fatih/color"
)

type Address struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run cep.go <cep>")
		return
	}

	cep := os.Args[1]

	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	address := Address{}

	err = json.Unmarshal(body, &address)
	if err != nil {
		panic(err)
	}

	fmt.Println()

	// Título com fundo azul e texto branco
	title := color.New(color.BgBlue, color.FgWhite, color.Bold)
	title.Println(" ENDEREÇO ENCONTRADO                                      ")

	// Criando cores personalizadas para cada campo
	labelColor := color.New(color.BgHiBlue, color.FgBlack, color.Bold)
	valueColor := color.New(color.BgHiBlack, color.FgWhite)
	bgColor := color.New(color.BgHiBlack)

	// Função auxiliar para imprimir campos alinhados
	printField := func(label, value string) {
		labelColor.Printf(" %-15s", label+":")  // -20 alinha à esquerda com 20 caracteres
		valueColor.Printf(" %-40s", value)      // alinha à esquerda com 80 caracteres
		bgColor.Print(" ")                      // preenche o resto da linha
		fmt.Println()
	}

	// Preenchendo a primeira linha com background
	bgColor.Println("                                                          ")

	// Imprimindo os campos com formatação
	printField("Cidade", address.Localidade)
	printField("UF", address.Uf)
	printField("Logradouro", address.Logradouro)
	printField("Bairro", address.Bairro)

	if address.Complemento != "" {
		printField("Complemento", address.Complemento)
	}

	printField("CEP", address.Cep)

	// Preenchendo a última linha com background
	bgColor.Println("                                                          ")
}
