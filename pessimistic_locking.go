package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var count int = 0

func doCount() {
    mu.Lock()
    count++
    mu.Unlock()
    wg.Done()
}

func main() {
    count = 0
    for i:= 0; i < 1000000; i++ {
        wg.Add(1)
        go doCount()
    }

    wg.Wait()
    fmt.Print(count)
}
