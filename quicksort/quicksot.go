package main

import "fmt"

func main() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	quicksort(arr, 0, len(arr) - 1)

	fmt.Println("Lista final ordenada: ", arr)
}

func quicksort(arr []int, inicio, fim int) {
	if(inicio >= fim) {
		return 
	}

	p := partition(arr, inicio, fim)

	quicksort(arr, inicio, p-1)
	quicksort(arr, p+1, fim)

}

func partition(arr []int, inicio, fim int) int {
	p := arr[fim]

	i := inicio - 1

	for j := inicio; j < fim; j++ {
		if arr[j] <= p {
			i++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}

	i++

	arr[i], arr[fim] = arr[fim], arr[i]

	return i
}