package main

import (
	"fmt"
	"sync"
)

// Atomic stores info about out atomic values
type Atomic struct {
	value int
	lock  sync.Mutex
}

// Increase increases the value from Atomic struct
func (i *Atomic) Increase() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

// Decrease decreases the value from Atomic struct
func (i *Atomic) Decrease() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value--
}

// Value returns the value from Atomic struct
func (i *Atomic) Value() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

var (
	counter = 0

	lock sync.Mutex

	atomicCounter = Atomic{}
)

func updateCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()

	counter++
	atomicCounter.Increase()

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}

	wg.Wait()

	fmt.Println(fmt.Sprintf("final counter: %d", counter))
	fmt.Println(fmt.Sprintf("final atomic counter value: %d", atomicCounter.Value()))
}
