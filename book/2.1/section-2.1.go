package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello from section 2.1")

	i, f, s := 10, 3.3, "Ola mundo"

	// isso NÃO funciona
	//var nome string, idade int, altura float64 = "João", 25, 1.75

	var (
		nome   string
		idade  int
		altura float64
	)

	nome, idade, altura = "João", 25, 1.75

	fmt.Printf("Nome: %s Idade %d Altura %f\n", nome, idade, altura)

	fmt.Println(i, " ", f, " ", s)

	fmt.Println("O Tipo da variável nome é: ", reflect.TypeOf(nome))
	fmt.Println("O Tipo da variável idade é: ", reflect.TypeOf(idade))
	fmt.Println("O Tipo da variável altura é: ", reflect.TypeOf(altura))

	str := "Hello"

	fmt.Println(str)

	bytesStr := []byte(str)

	fmt.Println(bytesStr)

	fmt.Println()

	for i, v := range bytesStr {
		fmt.Printf("indice %d tem o valor ascii %v\n", i, v)
	}

	fmt.Println("\nUsando array de rune")

	var str2 string = "WILLAMS"
	runes := []rune(str2)

	for i, r := range runes {
		fmt.Printf("Indice %d tem o valor ascii: %d, e o caractere %c\n", i, r, r)
	}

	fmt.Print(runes[2])

	scope()

}

func scope() {
	var x = "fora"

	fmt.Println()
	fmt.Println("Fora do bloco: ", x)

	func() {
		x := "dentro"
		fmt.Println("Dentro do bloco: ", x)
	}()

	fmt.Println("Fora do bloco novamente: ", x)

	// Exemplo com for
	x = "valor inicial"
	fmt.Println("\nAntes do for:", x)

	for i := 0; i < 3; i++ {
		x := fmt.Sprintf("valor no loop %d", i)
		fmt.Println("Dentro do for:", x)
	}

	fmt.Println("Depois do for:", x)

	// Exemplo com if
	x = "valor antes do if"
	fmt.Println("\nAntes do if:", x)

	if true {
		x := "valor dentro do if"
		fmt.Println("Dentro do if:", x)
	}

	fmt.Println("Depois do if:", x)

	// Exemplo com switch
	x = "valor antes do switch"
	fmt.Println("\nAntes do switch:", x)

	switch {
	case true:
		x := "valor dentro do switch"
		fmt.Println("Dentro do switch:", x)
	}

	fmt.Println("Depois do switch:", x)
}
