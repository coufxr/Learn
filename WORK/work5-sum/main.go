package main

import "math/rand"

type job struct {
	value int64
}
type result struct {
	job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)

func fun(f chan<- *job) {
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		f <- newJob
	}
}

func main() {

}
