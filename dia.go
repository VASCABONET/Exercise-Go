package main

import (
	"fmt"
	"strings"
)

func main() {
	var r string
	fmt.Println("Como esta el dia?")

	fmt.Println("Indica el dia segun la inicial del clima")
	fmt.Println("S: Soleado\nN:Nublado\nL:Lloviendo")
	fmt.Scanln(&r)
	R := (strings.ToUpper(r))
	switch R {
	case "S":
		fmt.Println("El dia esta Soleado")
	case "N":
		fmt.Println("El dia esta Nublado")
	case "L":
		fmt.Println("El dia esta Lluvioso")
	default:
		fmt.Println("No ha seleccionado nunguna de las letras")
	}

}
