package main

import (
	"fmt"
	"sync"
)

//goroutiune are used for multithreadin, concurrently running the task
//it does not ensure order of task
func task (id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main (){
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}