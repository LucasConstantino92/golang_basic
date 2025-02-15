package main

import "fmt"

func sum(n1 int, n2 int) int {
	return n1 + n2
}

func mathOperation(n1, n2 int16) (int16, int16) {
	sum := n1 + n2
	sub := n1 - n2

	return sum, sub
}

func main() {

	//Sum example
	sum1 := sum(10, 45)
	fmt.Println(sum1)

	//var function example
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	result := f("Function F")
	fmt.Println(result)

	//Multiple return example
	res1, res2 := mathOperation(10, 8)
	fmt.Println(res1, res2)
}
