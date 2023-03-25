package main

import "fmt"

var dinero = 100

type hijo struct {
	nombre      string
	sePortoBien bool
	mesada      int
}

func main() {

	hijos := []hijo{
		{
			nombre:      "emiliano",
			sePortoBien: true,
			mesada:      25,
		},
		{
			nombre:      "ronaldo",
			sePortoBien: false,
			mesada:      0,
		},
	}

	fmt.Println(hijos)

}
