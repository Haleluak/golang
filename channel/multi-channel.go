package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers = 2
	numTasks   = 7
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano()) // Tạo nhân gieo giá trị ngẫu nhiên
}

func main() {
	// Tạo kênh đa chứa các công việc
	tasks := make(chan string, numTasks)
	wg.Add(numWorkers)

	for gr := 1; gr <= numWorkers; gr++ {
		go worker(tasks, gr)
	}
	for t := 1; t <= numTasks; t++ {
		tasks <- fmt.Sprintf("công việc thứ %d.", t)
	}

	close(tasks)

	wg.Wait()
}
func worker(tasks chan string, worker int) {
	defer wg.Done()
	for {
		// Chờ nhận công việc
		task, ok := <-tasks

		if !ok {
			// Thắng nếu kênh đã đóng
			fmt.Printf("Người thứ %d: Hoàn thành các công việc được giao!\n", worker)
			return
		}
		fmt.Printf("Người thứ %d: Bắt đầu thực hiện %s\n", worker, task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Người thứ %d: Thực hiện xong %s\n", worker, task)
	}
}
