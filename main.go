package main

import (
	"fmt"
	"strconv"
)

func main() {
	aa := 122
	a := "9495"

	if aa > 12 {
		fmt.Println("Kewl!!")
		fmt.Println(a)
		var z int64 = 12
		fmt.Println(z)
	}

	if i, err := strconv.Atoi(a); err == nil {
		fmt.Println("Value of I", i)
		fmt.Println("Program ENDS", test(i, 2))
	} else {
		fmt.Println(err)
	}

}

func test(x int, y int) int {
	return x + y
}
