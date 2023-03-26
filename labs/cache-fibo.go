package labs

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Memory struct {
	f     Function
	cache map[int]FunctionResult
	lock  sync.Mutex
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	return &Memory{
		f:     f,
		cache: make(map[int]FunctionResult),
	}
}

func (m *Memory) Get(key int) (interface{}, error) {
	m.lock.Lock()
	result, exists := m.cache[key]
	m.lock.Unlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.f(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func StartCacheFibo() {
	cache := NewCache(GetFibonacci)
	fibo := []int{22, 33, 40, 42, 40, 42}

	var wg sync.WaitGroup

	for _, n := range fibo {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			start := time.Now()
			value, err := cache.Get(index)

			if err != nil {
				log.Println(err)
			}

			fmt.Printf("Fibo(%d) = %d -> %f seg\n", index, value, time.Since(start).Seconds())
		}(n)
	}
	wg.Wait()
}
