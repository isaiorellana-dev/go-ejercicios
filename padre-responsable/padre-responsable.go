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
	var hijosPortaronBien []hijo
	var hijosPortaronMal []hijo

	for i, value := range hijos {
		if value.nombre == hijoName {
			hijoIndex = i
		}
		if value.sePortoBien && value.nombre != hijoName {
			hijosPortaronBien = append(hijosPortaronBien, value)
		}
		if !value.sePortoBien {
			hijosPortaronMal = append(hijosPortaronMal, value)
		}
	}

	var dineroDeLosQueSePortaronBien float32

	for _, value := range hijosPortaronBien {
		dineroDeLosQueSePortaronBien += value.mesada
	}

	var dineroParaDarleAlHijo = dinero - dineroDeLosQueSePortaronBien

	if cantidad <= dineroParaDarleAlHijo {
		hijos[hijoIndex].mesada = cantidad

		var dineroARepartirALosPlebeyos = dinero - dineroDeLosQueSePortaronBien - cantidad

		for i := range hijos {
			if !hijos[i].sePortoBien {
				var hijosQueSePortanMalLen = len(hijosPortaronMal)
				hijos[i].mesada = dineroARepartirALosPlebeyos / float32(hijosQueSePortanMalLen)
			}
		}
		fmt.Println(hijos)
	} else {
		fmt.Println("No hay tanta plata")
	}
}

func main() {
	repartir("messi", 50)
}
