// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if x, y := rand.Int(), rand.Int(); x > y {
		fmt.Println("X is larger than Y")
	} else {
		fmt.Println("X is smaller than Y")
	}

	// ERROR: if文で宣言した変数は束縛されない
	// fmt.Println("X: " + x + "/ Y:" + y)

	// また、Go では if は式ではない
	//
	//  z := if x, y := rand.Int(), rand.Int(); x > y {
	//  	return "BIG"
	//  } else {
	//  	return "SMALL"
	//  }
	//
	//  fmt.Println(z)
}
