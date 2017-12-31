package main

import (
	"fmt"
	"strings"
	"sync"
)

const errorsCapacity = 1024

type errorReporter struct {
	errors    chan error
	counts    map[string]int
	waitGroup sync.WaitGroup
}

func newErrorReporter() *errorReporter {
	return &errorReporter{
		make(chan error, errorsCapacity),
		map[string]int{},
		sync.WaitGroup{},
	}
}

func (r *errorReporter) Errors() chan<- error {
	return r.errors
}

func (r *errorReporter) Analyze() {
	r.waitGroup.Add(1)

	for err := range r.errors {
		s := err.Error()

		if _, ok := r.counts[s]; !ok {
			r.counts[s] = 0
		}

		r.counts[s]++
	}

	r.waitGroup.Done()
}

func (r *errorReporter) Report() string {
	close(r.errors)
	r.waitGroup.Wait()

	if len(r.counts) == 0 {
		return ""
	}

	ss := make([]string, 0, len(r.counts))

	for s, i := range r.counts {
		ss = append(ss, fmt.Sprintf("%v\t%v", s, i))
	}

	return strings.Join(ss, "\n")
}
