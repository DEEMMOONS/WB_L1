package main

import "fmt"

func main() {
  in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  i := 5
  ans1 := append(in[:i], in[i+1:]...)
  fmt.Println(ans1) 
  ans2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
  ans2[i] = ans2[len(ans2) - 1]
  ans2[len(ans2) - 1] = 0
  ans2 = ans2[:len(ans2) - 1]
  fmt.Println(ans2)
}
