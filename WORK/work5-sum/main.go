package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type job struct {
	value int64
}
type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func fun(f chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		f <- newJob
	}
}
func f2(f <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-f
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go fun(jobChan)
	for i := 0; i < 24; i++ {
		go f2(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Println("value:", result.job.value, "sum:", result.result)
	}
	wg.Wait()
}
