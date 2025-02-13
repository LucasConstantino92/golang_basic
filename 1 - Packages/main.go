package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Writing from main")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("lconstantino@gmail.com")

	fmt.Println(erro)
}
