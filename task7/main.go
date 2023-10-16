package main

import (
  "sync"
  "fmt"
)

func main() {
  myMap1 := make(map[int]int)
  var wg sync.WaitGroup
  var mu sync.Mutex
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(i int, mu *sync.Mutex, wg *sync.WaitGroup) {
      defer wg.Done()
      mu.Lock()
      defer mu.Unlock()
      myMap1[i] = i*i
    }(i, &mu, &wg)
  }
  wg.Wait()
  fmt.Println(myMap1)

  // в пакете sync есть структура sync.Map обеспечивающая безопасную работу с мапами

  var myMap2 sync.Map
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(ind int) {
      defer wg.Done()
      myMap2.Store(ind, ind * ind)
    }(i)
  }
  wg.Wait()
  fmt.Printf("map[")
  for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(ind int) {
      defer wg.Done()
      value, _ := myMap2.Load(ind)
      fmt.Printf("%d:%d ", ind, value)
    }(i)
  }
  wg.Wait()
  fmt.Println("]")
}
