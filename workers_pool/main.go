package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Llongfile)
	var workChan = make(chan int)
	wg := sync.WaitGroup{}
	workersWg := sync.WaitGroup{}
	for i := 1; i <= 3; i++ {
		workersWg.Add(1)
		go doWork(i, workChan, &wg, &workersWg)
	}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		workChan <- i
	}
	wg.Wait()
	close(workChan)
	workersWg.Wait()
}

func doWork(i int, workChan chan int, wg, workersWg *sync.WaitGroup) {
	defer workersWg.Done()
	for w := range workChan {
		log.Println("doing work", w, "by worker", i)
		time.Sleep(time.Duration(rand.Intn(i) * 100))
		wg.Done()
	}
	log.Println("worker", i, "finished his job")
}
