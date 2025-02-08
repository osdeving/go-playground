package main

import "fmt"

func main() {

	arr := [5]int{}

	langs := [3]string{"Go", "Python", "Rust"}

	inteiros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	inteiros[4] = 42

	decimais := [5]float64{}

	letras := [3]rune{'a', 'b', 'c'}

	fmt.Printf("%c%c%c\n", letras[0], letras[1], letras[2])

	decimais[0] = 3.14
	decimais[1] = 6.28
	decimais[2] = 1.618
	decimais[3] = 2.718
	decimais[4] = 0.577

	fmt.Println(decimais)

	fmt.Println(inteiros)

	fmt.Println(langs)

	fmt.Println(arr)

	inteiros2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Printf("Array original %v array invertido %v\n", inteiros2, inverterArray(inteiros2))

	fmt.Printf("%d\n", len(inteiros2))

	fmt.Println(inteiros2)


	diasSemana()
}

func diasSemana() {
	dias := [7]string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}

	fmt.Print("|")

	for i, dia := range dias {
		fmt.Printf(" Dia %d: %s |", i, dia)
	}
}	

func inverterArray(arr[6]int) [6]int {
	for i := 0; i < len(arr) / 2; i++ {
		j := len(arr) - 1 - i
		arr[j], arr[i] = arr[i], arr[j]		
	}	

	return arr
}