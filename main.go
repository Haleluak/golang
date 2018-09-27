package main

import (
	"fmt"
	"sync/atomic"
)

var counter int64

func main() {
	fmt.Println("Bat dau GoRoutines")
	for i := 0; i < 100; i++ {
		go func() { // Hàm vô danh
			for i := 0; i < 10000; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	fmt.Println("Doi ket thuc GoRoutines")
	var input string
	fmt.Scanln(&input)
	fmt.Println(counter)
}
