package variables

import "fmt"

func Println() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	interpolation()
}

func interpolation() {
	name := "Alice"
	age := 25

	// Menggunakan operator + untuk menggabungkan string
	result := "Her name is " + name + " and she is " + fmt.Sprint(age) + " years old."
	fmt.Println(result)
}
