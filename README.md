# Go Concurrency Issue: Data Race

This repository demonstrates a common concurrency issue in Go programs: data races.  Data races occur when multiple goroutines access and modify the same shared variable without proper synchronization mechanisms, leading to unpredictable and incorrect results.

The `bug.go` file contains the buggy code, showcasing the data race. The `bugSolution.go` file provides a corrected version using mutexes to prevent data races.

## How to Reproduce

1. Clone this repository.
2. Navigate to the project directory.
3. Run `go run bug.go`. You'll likely see inconsistent or unexpected output.
4. Run `go run bugSolution.go`. This demonstrates the corrected version with consistent output.

## Solution

The solution involves using synchronization primitives, such as mutexes, to protect shared resources from concurrent access. This ensures that only one goroutine can access and modify the critical section at a time, preventing data races.