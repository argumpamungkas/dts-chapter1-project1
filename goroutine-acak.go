package main

import (
	"fmt"
	"sync"
)

func main() {

	bisa := []interface{}{"bisa1", "bisa2", "bisa3"}
	coba := []interface{}{"coba1", "coba2", "coba3"}

	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go print(i, bisa, &wg)
	}

	for j := 1; j <= 4; j++ {
		wg.Add(1)
		go print(j, coba, &wg)
	}

	wg.Wait()
}

func print(num int, data interface{}, wg *sync.WaitGroup) {
	fmt.Printf("%v %d\n", data, num)
	wg.Done()
}
