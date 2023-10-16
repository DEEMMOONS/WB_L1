package main

import "fmt"

func main() {
  strs := []string{"cat", "cat", "dog", "cat", "tree"}
  set := make(map[string]bool)
  for _, item := range strs {
    set[item] = true
  }
  fmt.Println(set)
}
