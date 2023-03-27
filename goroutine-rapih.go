package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	var mtx sync.Mutex

	coba := []interface{}{"coba1", "coba2", "coba3"}
	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go func(num int) {
			mtx.Lock()
			if num%2 == 1 {
				fmt.Println(coba, num)
			} else {
				fmt.Println(bisa, num)
			}
			mtx.Unlock()
			wg.Done()
		}(i)
	}

	time.Sleep(time.Microsecond)

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go func(num int) {
			mtx.Lock()
			if num%2 == 1 {
				fmt.Println(bisa, num)
			} else {
				fmt.Println(coba, num)
			}
			mtx.Unlock()
			wg.Done()
		}(j)
	}

	wg.Wait()

}
