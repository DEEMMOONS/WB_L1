package main

import "fmt"

func main() {
  var in, out string = "", ""
  fmt.Scan(&in)
  tmp := []rune(in)
  for i := len(tmp) - 1; i >= 0; i-- {
    out += string(tmp[i])
  }
  fmt.Println(out)
}
