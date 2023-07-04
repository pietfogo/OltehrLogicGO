// Other logic
package main

import (
	"fmt"
	"strings"
)

func main()  {
	var texto string
	var caractere string
	var vezes string
	fmt.Print("Escreva uma palavra ou texto: ")
	fmt.Scan(&texto)
	fmt.Print("\nDigite o caractere a ser contado: ")
	fmt.Scan(&caractere)
	contador := strings.Count(texto, caractere)
	if contador > 1 {
		vezes= "vezes";
	}else if contador == 0 {
		vezes= "vez";
	} else {
		vezes= "No caso nenhuma vez";
	}

	fmt.Printf("O caractere '%v' aparece '%v' %v no texto '%v'", caractere, contador, vezes, texto)
}