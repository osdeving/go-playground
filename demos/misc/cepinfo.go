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

type ViaCep struct {
	Rua string `json:"logradouro"`
	Bairro string `json:"bairro"`
	Cidade string `json:"localidade"`
	Estado string `json:"estado"`
	Cep string `json:"cep"`
}

func (c *ViaCep) Print() {
	title := color.New(color.BgBlue, color.FgHiWhite, color.BlinkRapid)
	
	title.Print("                   ENDEREÇO ENCONTRADO                      ")
	title.Println()

	printField := func(label, value string) {
		labelColor := color.New(color.BgCyan, color.FgHiWhite)
		labelColor.Printf("⏩ %-15s", label)

		valueColor := color.New(color.BgGreen, color.FgBlack)
		valueColor.Printf(" %-40s ", value)
		valueColor.Println()
	}

	printField("Rua", c.Rua)
	printField("Bairro", c.Bairro)
	printField("Cidade", c.Cidade)
	printField("Estado", c.Estado)
	printField("CEP", c.Cep)
	
	color.New(color.BgBlue).Println("                                                            ")

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./cepinfo <cep>")
		return
	}

	cep := os.Args[1]

	endpoint := "https://viacep.com.br/ws/" + cep + "/json";
	
	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	viacep := ViaCep{}

	bytes, err := io.ReadAll(res.Body)
	
	json.Unmarshal(bytes, &viacep)
	
	viacep.Print()
}