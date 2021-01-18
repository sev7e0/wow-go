package main

import (
	"fmt"
	"math"
)

const (
	CC string = "CC"
)

var (
	A string = "test"
	B bool = false
	)

func main() {
	fmt.Printf("print test :%g\n", math.Pi)
	fmt.Println(add(32,32))
	fmt.Println(swap("world", "hello"))
	c, d := split(30)
	fmt.Println(c,d)

}

func add(x, y int) int {
	return x*y
}

func swap(a, b string) (string, string) {
	return b, a
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}