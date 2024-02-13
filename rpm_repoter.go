package main

import (
	"fmt"
	"sync"
	"time"
)

const successesCapacity = 1024

type rpmReporter struct {
	successes chan bool
	count     int
	minutes   float64
	waitGroup sync.WaitGroup
}

func newRPMReporter() *rpmReporter {
	return &rpmReporter{
		make(chan bool, successesCapacity),
		0,
		0,
		sync.WaitGroup{},
	}
}

func (r *rpmReporter) Successes() chan<- bool {
	return r.successes
}

func (r *rpmReporter) Analyze() {
	r.waitGroup.Add(1)

	t := time.Now()

	for range r.successes {
		r.count++
	}

	r.minutes = time.Since(t).Minutes()

	r.waitGroup.Done()
}

func (r *rpmReporter) Report() string {
	close(r.successes)
	r.waitGroup.Wait()

	return fmt.Sprint(float64(r.count) / r.minutes)
}
