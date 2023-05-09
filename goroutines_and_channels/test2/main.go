package main

import (
    "fmt"
    "runtime"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        fib := fibonacci(j)
        fmt.Printf("Worker %d finished job %d, result: %d\n", id, j, fib)
        results <- fib
    }
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    return fibonacci(n - 1) + fibonacci(n - 2)
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= runtime.NumCPU(); w++ {
        go worker(w, jobs, results)
    }

    for j := 1; j <= 10; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= 10; a++ {
        <-results
    }
}
