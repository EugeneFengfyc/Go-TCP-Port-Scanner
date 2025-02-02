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
	sem := make(chan struct{}, 100)    // 限制最大并发数量为 100
	var openPorts atomic.Value         // 用 atomic 包存储并发安全的端口列表
	openPorts.Store(make([]string, 0)) // 初始化为一个空切片

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(j int) {
			defer wg.Done()
			defer func() { <-sem }()

			address := fmt.Sprintf("192.168.1.88:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return // 端口关闭，忽略
			}
			conn.Close()
			// 使用原子操作确保并发安全
			ports := openPorts.Load().([]string)
			ports = append(ports, address)
			openPorts.Store(ports)
		}(i)
	}
	wg.Wait()

	elapsed := time.Since(start) / 1e9
	// 最后统一打印开放的端口
	fmt.Printf("Open ports:\n")
	for _, port := range openPorts.Load().([]string) {
		fmt.Println(port)
	}

	fmt.Printf("\n\n%d seconds\n", elapsed)
}
