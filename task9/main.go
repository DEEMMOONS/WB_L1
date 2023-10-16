package main

import (
  "fmt"
  "sync"
)

func Work(wg *sync.WaitGroup, in, out chan int) {
  defer wg.Done()
  for i := range in {
    out <- i * 2
  }
}

func main() {
  var num int
  fmt.Scanf("%d",&num)

	arr := [10]int {}
	for i := 0; i < 10; i++ {
		arr[i] = i
	}

  var wg, wg2 sync.WaitGroup
  wg.Add(1)
  wg2.Add(1)
  chIn := make(chan int)
	chOut := make(chan int)

  for i := 1; i <= num; i++ {
		wg.Add(1)
		go Work(&wg, chIn, chOut)
	}

	go func() {
    defer wg.Done()
    for i := range arr {
      chIn <- i
    }
    close(chIn)
    }()

	go func() {
    defer wg2.Done()
		for i := range chOut{
			fmt.Println(i)
		}
  }()
  wg.Wait()
	close(chOut)
  wg2.Wait()
}
