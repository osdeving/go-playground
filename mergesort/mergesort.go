package main

import "fmt"

func main() {
	fmt.Println("Lista final ordenada: ", mergesort([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
}

func merge(arr1 []int, arr2 []int) []int {

	arr := make([]int, len(arr1) + len(arr2))
	
	idx, idx1, idx2 := 0, 0, 0

	for idx1 < len(arr1) && idx2 < len(arr2) {
		if(arr1[idx1] < arr2[idx2]) {
			arr[idx] = arr1[idx1]
			idx1++
		} else {
			arr[idx] = arr2[idx2]
			idx2++
		}		
		idx++
	}

	copy(arr[idx:], arr1[idx1:])
    copy(arr[idx:], arr2[idx2:])

	return arr
}

func mergesort(arr []int) []int {
	
	if(len(arr) <= 1) {
		return arr
	}

	meio := len(arr) / 2

	esquerda := mergesort(arr[:meio])
	direita := mergesort(arr[meio:])

	return merge(esquerda, direita)
}