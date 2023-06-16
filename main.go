package main

import (
	"fmt"
	"runtime"
	"time"
)

func allocateMemory() {
	var data [][]int

	for i := 0; i < 10000; i++ {
		// Create a slice that uses more memory every time.
		data = append(data, make([]int, 1000000))
	}

	fmt.Println("Memory allocated successfully.")
}

func main() {
	// Prints the amount of memory before allocating.
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Initial Memory Usage (bytes):", m.Alloc)

	// allocate memory in function allocateMemory
	allocateMemory()
	allocateMemory()

	// Prints the amount of memory after allocation.
	runtime.ReadMemStats(&m)
	fmt.Println("Final Memory Usage (bytes):", m.Alloc)

	// Wait a moment for Memory Usage to be checked in pprof.
	time.Sleep(30 * time.Second)
}
