package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
    queue []int32
    mutex sync.Mutex
}

func (q *ConcurrentQueue) Enqueue(item int32) {
    q.mutex.Lock()
    defer q.mutex.Unlock()
    q.queue = append(q.queue,item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
  q.mutex.Lock()
  defer q.mutex.Unlock()
    if(len(q.queue) == 0) {
        panic("queue is empty")
    }
    item := q.queue[0]
    q.queue = q.queue[1:]
    return item
}

func (q *ConcurrentQueue) Size() int {
    return len(q.queue)
}

var wg sync.WaitGroup
func main() {
  q1 := ConcurrentQueue {
      queue: make([]int32,0),
  }

  
  for i:=0; i < 1e6; i++ {
    wg.Add(1)
    go func() {
        q1.Enqueue(rand.Int31())
        wg.Done()
    }()
  }

  wg.Wait()
  fmt.Println(q1.Size())

  for i:=0; i < 1e6; i++ {
    wg.Add(1)
    go func() {
        q1.Dequeue()
        wg.Done()
    }()
  }

  wg.Wait()
  fmt.Println(q1.Size())
  /*
  q1.Enqueue(1)
  q1.Enqueue(2)
  q1.Enqueue(3)
  fmt.Println(q1.Dequeue())
  fmt.Println(q1.Dequeue())
  fmt.Println(q1.Dequeue())
  fmt.Println(q1.Dequeue())
	fmt.Println("Hello, World!")
*/
}
