package main

import (
  "fmt"
  "os"
  "strconv"
)

var last, now int // define global integer var

// Main function
func main() {

  last = 0
  now  = 1
  
  s      := os.Args[1:]        // Get user paramn
  max, _ := strconv.Atoi(s[0]) // Conert string to integer
  
  for i := 1; i < max; i++ {
    fmt.Printf("%d\n",Fibonacci()) // Call fibonacci function
  }

}

// Fibonacci function return the next value
func Fibonacci() int {
  
  total := last + now
  now   = last
  last  = total
  
  return total
}