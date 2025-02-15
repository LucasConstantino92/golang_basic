package main

import (
	"errors"
	"fmt"
)

func main() {
	//int, int8, int16, int32, int64
	var number int = 10000000000000
	fmt.Println(number)

	//uint - Unregisted int
	var number2 uint = 100000
	fmt.Println(number2)

	//alias
	//int32 = rune
	var number3 rune = 1000000
	fmt.Println(number3)

	//byte = uint8
	var number4 byte = 123
	fmt.Println(number4)

	//float32, float64
	var realNumber float32 = 123.321
	fmt.Println(realNumber)

	//String
	var str string = "String"
	fmt.Println(str)

	//Char
	//Always return ascii code
	var char rune = 'A'
	fmt.Println(char)

	//Boolean
	var isTrue bool = false
	fmt.Printf("%t", isTrue)

	//Error
	var err error = errors.New("Internal Error")
	fmt.Println(err)
}
