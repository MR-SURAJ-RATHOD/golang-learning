package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID    int
	Value int
}

type Result struct {
	JobID  int
	Output int
	Err    error
}

func main() {
	fmt.Println("=== 03 Concurrency & Context ===")

	lesson01Goroutines()
	lesson02Mutex()
	lesson03WorkerPool()
}

// --- Lesson 1: Goroutines & WaitGroup ---
func lesson01Goroutines() {
	fmt.Println("\n--- Lesson 1: Goroutines & WaitGroup ---")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  Goroutine %d done\n", id)
		}(i)
	}
	wg.Wait()
	fmt.Println("  All goroutines finished.")
}

// --- Lesson 2: Mutex (protect shared state) ---
func lesson02Mutex() {
	fmt.Println("\n--- Lesson 2: Mutex ---")

	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("  Counter (with mutex): %d\n", counter)
}

// --- Lesson 3: Worker Pool + Context Timeout ---
func lesson03WorkerPool() {
	fmt.Println("\n--- Lesson 3: Worker Pool + Context ---")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	jobs := make(chan Job, 10)
	results := make(chan Result, 10)

	var wg sync.WaitGroup
	numWorkers := 3

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, &wg, jobs, results)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- Job{ID: i, Value: rand.Intn(100)}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Printf("  ❌ Job %d failed: %v\n", res.JobID, res.Err)
		} else {
			fmt.Printf("  ✅ Job %d → output %d\n", res.JobID, res.Output)
		}
	}
	fmt.Println("  Worker pool done.")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup, jobs <-chan Job, results chan<- Result) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("  ⚠️ Worker %d stopped: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			select {
			case <-time.After(time.Duration(rand.Intn(400)) * time.Millisecond):
				results <- Result{JobID: job.ID, Output: job.Value * 2}
			case <-ctx.Done():
				results <- Result{JobID: job.ID, Err: ctx.Err()}
			}
		}
	}
}
