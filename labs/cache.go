package labs

import (
	"fmt"
	"sync"
	"time"
)

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Lock       sync.RWMutex
}

func LongProcess(n int) int {
	fmt.Printf("Calculate LongProcess(%d)\n", n)
	time.Sleep(5 * time.Second)
	return n
}

func (s *Service) Work(job int) {
	s.Lock.RLock()
	exists := s.InProgress[job]
	if exists {
		s.Lock.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Lock.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Lock.Unlock()

		fmt.Printf("Waiting for response job: %d\n", job)
		resp := <-response
		fmt.Printf("Response done, received: %d\n", resp)

		return
	}

	s.Lock.RUnlock()

	s.Lock.Lock()
	s.InProgress[job] = true
	s.Lock.Unlock()

	fmt.Printf("Calculate LongProcess(%d)\n", job)
	result := LongProcess(job)

	s.Lock.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Lock.RUnlock()

	if exists {
		for _, pependingWorker := range pendingWorkers {
			pependingWorker <- result
		}
		fmt.Printf("Result sent all pending workers job %d\n", job)
	}

	s.Lock.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Lock.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func StartCache() {
	service := NewService()
	jobs := []int{3, 4, 5, 4, 5, 8, 8, 8}
	var wg sync.WaitGroup
	wg.Add(len(jobs))

	for _, n := range jobs {
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(n)
	}
	wg.Wait()
}
