package main

import (
  "fmt"
  "strings"
)

func Uniq(in string) bool {
  runes := make(map[rune]int)
  for _, char := range in {
    runes[char]++
  }
  for _, count := range runes {
    if count > 1 {
      return false
    }
  }
  return true
}

func main() {
  var in string
  fmt.Scan(&in)
  in = strings.ToLower(in)
  if Uniq(in) {
    fmt.Println("true")
  } else {
    fmt.Println("false")
  }
}
