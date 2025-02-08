//go:build ignore

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var x uint = 0b011000011
	count := bits.OnesCount(x)

	fmt.Printf("count: %v\n", count)
}
