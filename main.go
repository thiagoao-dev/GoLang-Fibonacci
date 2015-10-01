package main

import "fmt"

var last int
var now  int

func main() {

  last = 0
  now  = 1
  
  for i := 1; i < 10; i++{
    fmt.Printf("%d\n",fibonacci())    
  }

}

func fibonacci() int {
  
  total := last + now
  now = last
  last = total
  
  return total
}