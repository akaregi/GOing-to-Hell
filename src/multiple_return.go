// +build ignore

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x, y := tuple()

	fmt.Println(x)
	fmt.Println(y)
}

func tuple() (int, int) {
	x := rand.Intn(10)
	y := rand.Intn(10)

	return x, y
}
