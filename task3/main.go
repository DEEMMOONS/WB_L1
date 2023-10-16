package main

import (
  "fmt"
  "sync"
)

func Sum (wg *sync.WaitGroup, res chan int) {
  defer wg.Done()
  sum := 0
  for n := range res {
    sum += n
  }
  fmt.Println(sum)
}

func main() {
  nums := []int{2, 4, 6, 8, 10}
  res := make(chan int)
  var wg, wg2 sync.WaitGroup
  wg.Add(len(nums))
  wg2.Add(1)
  go Sum(&wg2, res)
  for _, n := range nums {
    go func (n int) {
      defer wg.Done()
      res <- n * n
    }(n)
  }
  wg.Wait()
  close(res)
  wg2.Wait()
}
