// +build ignore

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 乱数生成な
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
}
