package main //defines executable package

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Hello, world")
	auxiliar.Write()

	erro := checkmail.ValidateFormat("test@mail.com")
	fmt.Println(erro)
}
