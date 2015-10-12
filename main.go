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
	for i := 0; i < rand.Intn(20); i++ {
		fmt.Println(i, f())
	}
}
