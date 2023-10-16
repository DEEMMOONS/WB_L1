package main

import (
  "fmt"
  "time"
)

func Read(ch chan int) {
  for val := range ch {
    fmt.Printf("Read value = %d\n", val)
  }
}

func Write(ch chan int) {
  for i := 0; ; i++ {
    ch <- i
    time.Sleep(time.Second)
  }
}

func main() {
  var n int
  fmt.Scan(&n)
  ch := make(chan int)
  defer close(ch)
  go Write(ch)
  go Read(ch)
  time.Sleep(time.Second * time.Duration(n))
  fmt.Println("End")
}
