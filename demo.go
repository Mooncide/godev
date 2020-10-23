package main

import (
	"log"
)

var c string = ""
var r int8 = 0

func demo() {

	a := 1
	b := 2
	c := 0

	log.Println(a, b, c)
	log.Println(fourComp(a, b, &c))
	log.Println(a, b, c)
	log.Println(sumall(a, b, c))

}
func Demo(x, y int) {

	a := x
	b := y
	c := 0

	log.Println(a, b, c)
	log.Println(fourComp(a, b, &c))
	log.Println(a, b, c)
	log.Println(sumall(a, b, c))

}
func fourComp(a int, b int, c *int) (r1, r2, r3, r4 int) {
	r1 = a + b
	r2 = a - b
	r3 = a * b
	r4 = a / b

	*c = r1 + r2 + r3 + r4

	return
}

func sumall(a ...int) (x int) {
	x = 0

	for _, v := range a {
		x += v
	}
	return
}
