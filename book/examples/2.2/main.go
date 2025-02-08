package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {
	fmt.Println(" 🔧🪛     Hello from section 2.2 🔧🪛")

	var f float64

	fmt.Print("🟢   Digite um número decimal: ")
	fmt.Scan(&f)

	i := int(f)

	fmt.Println(" ➡️ ", i)

	str := strconv.Itoa(i)

	fmt.Println(str)

	newI, _ := strconv.ParseInt(str, 0, 0)

	fmt.Printf("%d + %d = %d", i, newI, int64(i)+newI)

	str2 := "willams"
	var str3 []rune

	for _, v := range str2 {
		str3 = append(str3, unicode.ToUpper(v))
	}

	fmt.Println(string(str3))

	// Forma 1: declaração em bloco
	var (
		name     string
		age      int
		height   float64
		isActive bool
	)

	// Forma 2: na mesma linha
	var name, age, height, active = "John", 25, 1.75, true

}
