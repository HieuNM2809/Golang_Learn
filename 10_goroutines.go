package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// Danh sách các khoảng thời gian (sleep) khác nhau cho mỗi tác vụ
	sleepDurations := []time.Duration{
		2 * time.Second,
		1 * time.Second,
		3 * time.Second,
		1 * time.Second,
		2 * time.Second,
	}

	// Ghi lại thời gian bắt đầu của toàn bộ quá trình
	startTime := time.Now()

	for i, duration := range sleepDurations {
		wg.Add(1)
		go func(i int, duration time.Duration) {
			defer wg.Done()
			fmt.Printf("Goroutine %d bắt đầu, sleep %v...\n", i, duration)
			time.Sleep(duration)
			fmt.Printf("Goroutine %d kết thúc\n", i)
		}(i, duration)
	}

	// Chờ cho tất cả các goroutine hoàn thành
	wg.Wait()
	totalDuration := time.Since(startTime)
	fmt.Printf("Tổng thời gian xử lý: %v\n", totalDuration)
}
