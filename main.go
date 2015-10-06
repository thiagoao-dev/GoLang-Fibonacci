package main

import (
  "fmt"
  "github.com/thiagoaugustus/GoLang-Fibonacci/alg"
)

func main() {
	f := alg.fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}