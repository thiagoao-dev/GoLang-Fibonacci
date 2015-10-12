package main

import (
	"fmt"
	"github.com/thiagoaugustus/GoLang-Fibonacci/algorithm"
	"math/rand"
	"time"
)

func main() {
	f := algorithm.Fibonacci()
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(20)
	for i := 1; i <= x; i++ {
		fmt.Println(i, f())
	}
}
