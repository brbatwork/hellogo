package main

import (
  "runtime"
  "sync"
  "log"
)

// global vars
var (
  count int
  wg    sync.WaitGroup
  mutex sync.Mutex
)

func main() {
  runtime.GOMAXPROCS(runtime.NumCPU())
  wg.Add(2)

  go increment()
  go increment()

  wg.Wait()
  log.Printf("count: %d\n", count)
}

func increment() {
  for i := 0; i < 2; i++ {
    mutex.Lock()
    c := count
    runtime.Gosched() // NEVER DO THIS FOR EXAMPLE ONLY - Gives control back to the scheduler
    c++
    count = c
    mutex.Unlock()
  }

  wg.Done()
}
