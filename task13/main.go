package main

import "fmt"

func main() {
  var a, b int = 2, 5
  fmt.Println(a, b)
  a, b = b, a
  fmt.Println(a, b)
}
