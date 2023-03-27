package main

import (
	"fmt"
)

var dinero float32 = 100

type hijo struct {
	nombre      string
	sePortoBien bool
	mesada      float32
}

var hijos = []hijo{
	{
		nombre:      "emiliano",
		sePortoBien: false,
		mesada:      25,
	},
	{
		nombre:      "ronaldo",
		sePortoBien: false,
		mesada:      0,
	},
	{
		nombre:      "messi",
		sePortoBien: true,
		mesada:      10,
	},
	{
		nombre:      "dua lipa",
		sePortoBien: true,
		mesada:      25,
	},
}

func repartir(hijoName string, cantidad float32) {
	var hijoIndex int
	var hijosQueSePortaronBien []hijo
	var hijosQueSePortaronMal []hijo

	for i, value := range hijos {
		if value.nombre == hijoName {
			hijoIndex = i
		}
	}

	for i, value := range hijos {
		if value.sePortoBien && value.nombre != hijoName {
			hijosQueSePortaronBien = append(hijosQueSePortaronBien, hijos[i])
		}
	}

	for i, value := range hijos {
		if !value.sePortoBien {
			hijosQueSePortaronMal = append(hijosQueSePortaronMal, hijos[i])
		}
	}

	var dineroDeLosQueSePortaronBien float32

	for _, value := range hijosQueSePortaronBien {
		dineroDeLosQueSePortaronBien = dineroDeLosQueSePortaronBien + value.mesada
	}

	var dineroParaDarleAlHijo = dinero - dineroDeLosQueSePortaronBien

	if cantidad <= dineroParaDarleAlHijo {
		hijos[hijoIndex].mesada = cantidad

		var dineroParaRepartirALosPlebeyos = dinero - dineroDeLosQueSePortaronBien - cantidad

		for i := range hijos {
			if !hijos[i].sePortoBien {
				var hijosQueSePortanMalLen = len(hijosQueSePortaronMal)
				hijos[i].mesada = dineroParaRepartirALosPlebeyos / float32(hijosQueSePortanMalLen)
			}
		}
		fmt.Println(hijos)
	}
}

func main() {
	repartir("messi", 50)
}
