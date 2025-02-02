package main

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	sem := make(chan struct{}, 100)    // Limit the maximum concurrency to 100
	var openPorts atomic.Value         // Use the atomic package to store a thread-safe list of open ports
	openPorts.Store(make([]string, 0)) // Initialize as an empty slice

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int) {
			defer wg.Done()
			defer func() { <-sem }() // Release the semaphore after the goroutine finishes

			address := fmt.Sprintf("192.168.1.88:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return // Port is closed, ignore it
			}
			conn.Close()
			// Use atomic operations to ensure thread safety
			ports := openPorts.Load().([]string)
			ports = append(ports, address)
			openPorts.Store(ports)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start) / 1e9
	// Print the open ports at the end
	fmt.Printf("Open ports:\n")
	for _, port := range openPorts.Load().([]string) {
		fmt.Println(port)
	}

	fmt.Printf("\n\n%d seconds\n", elapsed)
}
