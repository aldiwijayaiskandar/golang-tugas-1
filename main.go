package main

import (
	"fmt"
	"tugas1/matematika"
)

func main() {
	result := matematika.CekGanjilGenap(1)
	result2 := matematika.CekGanjilGenap(1, 2, 3, 4, 5)

	fmt.Println("v1: ", result)
	fmt.Println("v2: ", result2)
}
