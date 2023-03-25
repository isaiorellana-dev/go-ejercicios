package main

import "fmt"

type decimal float32

type user struct {
	id     int
	name   string
	active bool
}

var x decimal
var y string
var z bool

var usuario user

func main() {

	x = 12.2
	y = "Hola, que pedo"
	z = true

	usuario = user{
		id:     9,
		name:   "Isaí",
		active: true,
	}

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)

	fmt.Println(usuario)

	fmt.Println(s)
}
