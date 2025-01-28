package main

import "fmt"

func main() {

	arr := [5]int{}

	langs := [3]string{"Go", "Python", "Rust"}

	inteiros := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	inteiros[4] = 42

	decimais := [5]float64{}

	letras := [3]rune{'a', 'b', 'c'}

	fmt.Println(letras)

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

	for i := 0; i < len(inteiros2) / 2; i++ {
		j := len(inteiros2) - i - 1
		inteiros2[i], inteiros2[j] = inteiros2[j], inteiros2[i]
	}

	fmt.Printf("%d\n", len(inteiros2))

	fmt.Println(inteiros2)


}