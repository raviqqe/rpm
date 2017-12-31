package main

import (
	"math"
	"sync"

	"github.com/valyala/fasthttp"
)

type attacker struct {
	fasthttp.Client
}

func newAttacker() attacker {
	return attacker{fasthttp.Client{MaxConnsPerHost: math.MaxInt32}}
}

func (a *attacker) Attack(url string, n int, c int, ss chan<- bool, es chan<- error) {
	q := make(chan bool, c)
	g := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		q <- true
		g.Add(1)

		go func() {
			_, _, err := a.Get(nil, url)

			if err == nil {
				ss <- true
			} else {
				es <- err
			}

			<-q
			g.Done()
		}()
	}

	g.Wait()
}
