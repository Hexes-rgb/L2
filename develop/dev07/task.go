package main

import (
	"fmt"
	"time"
)

func main() {
	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			time.Sleep(after)
			close(c)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Minute),
		sig(6*time.Minute),
		sig(7*time.Minute),
		sig(8*time.Minute),
		sig(1*time.Second),
		sig(1*time.Minute),
	)

	fmt.Printf("done after %v\n", time.Since(start))
}

func or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	default:
		orDone := make(chan struct{})
		go func() {
			defer close(orDone)
			switch len(channels) {
			case 2:
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...):
				}
			}
		}()
		return orDone
	}
}
