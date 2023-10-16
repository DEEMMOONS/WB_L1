package main

import (
  "sync"
  "fmt"
  "os"
  "syscall"
  "os/signal"
  "time"
)

func Work(wg *sync.WaitGroup, tasks chan int, ind int) {
  defer wg.Done()
  for task := range tasks {
    fmt.Printf("Worker %d read num = %d\n", ind, task)
  }
}

func main() {
  var n int
  fmt.Scan(&n)
  var wg sync.WaitGroup
  tasks := make(chan int, n)
  interruptChannel := make(chan os.Signal, 1)
	defer close(interruptChannel)
  signal.Notify(interruptChannel, syscall.SIGINT, syscall.SIGTERM)
  for i := 1; i <= n; i++ {
		wg.Add(1)
		go Work(&wg, tasks, i)
	}
  go func() {
    for i := 0; ; i++ {
        tasks <- i
        time.Sleep(time.Second)
    }
  }()
  <-interruptChannel
  close(tasks)
  wg.Wait()
}
