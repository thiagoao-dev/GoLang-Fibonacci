package main

import (
  "fmt"
  "github.com/thiagoaugustus/GoLang-Fibonacci/algorithm"
)

func main() {
	f := algorithm.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}