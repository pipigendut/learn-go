package main

import (
	"fmt"
	"strconv"
)

func split(sum int) (x int, y int, z string) {
	x = sum * 4 / 9
	y = sum - x

	var q string = "0"
	q = strconv.Itoa(sum)
	z = q
	fmt.Println(q)

	a := q
	fmt.Println(a)
	return
}

func main() {
	x, y, z := split(12)
	fmt.Println(x, y, z)
}
