package main

import "fmt"

func main() {
	var a string = "string var \n"
	b := "string too \n"

	fmt.Println(a, b)

	var (
		c string = "also string \n"
		d        = "and here too \n"
	)

	fmt.Println(c, d)

	f, g := "another var \n", "and another \n"

	fmt.Println(f, g)

	const constVar string = "const var \n"

	fmt.Println(constVar)

	c, d = f, g

	fmt.Println(c, d)

}
