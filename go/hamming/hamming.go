package hamming

import (
	"errors"
	"runtime"
	"sync"
)

//pair is a struct with two runes used for comparision.
type pair struct {
	a rune
	b rune
}

// preparePairs receives two strings and creates a channel for the pairs and in a goroutine pass the pairs for comparision(fan-out pattern).
func preparePairs(str1, str2 string, quit <-chan bool) <-chan pair {
	pairs := make(chan pair)
	go func() {
		for i := range str1 {
			select {
			case <-quit:
				return
			case pairs <- pair{rune(str1[i]), rune(str2[i])}:
			}
		}
		close(pairs)
	}()
	return pairs
}

//checkPair creates a bool channel and a goroutine to evaluate pairs received on the pairs channel and send the answer on the returned channel.
func checkPair(pairs <-chan pair, quit <-chan bool) <-chan bool {
	equal := make(chan bool)
	go func() {
		for pair := range pairs {
			select {
			case <-quit:
				return
			case equal <- pair.a == pair.b:
			}
		}
		close(equal)
	}()
	return equal
}

//merge joins all the channels in one(fan-in pattern).
func merge(quit <-chan bool, channels ...<-chan bool) <-chan bool {
	var wg sync.WaitGroup

	wg.Add(len(channels))
	outgoingCheckedPairs := make(chan bool)
	multiplex := func(c <-chan bool) {
		defer wg.Done()
		for i := range c {
			select {
			case <-quit:
				return
			case outgoingCheckedPairs <- i:
			}
		}
	}
	for _, c := range channels {
		go multiplex(c)
	}
	go func() {
		wg.Wait()
		close(outgoingCheckedPairs)
	}()
	return outgoingCheckedPairs
}

//Distance receives two strings and compare them returning the total of different pairs. It's optimized for parallel work and uses the concurrency patterns fan-in and fan-out.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, errors.New("strings have different size")
	}

	quit := make(chan bool)
	defer close(quit)

	pairs := preparePairs(a, b, quit)

	cores := runtime.NumCPU()
	routines := make([]<-chan bool, cores)
	for i := 0; i < cores; i++ {
		routines[i] = checkPair(pairs, quit)
	}

	distance := 0
	for i := range merge(quit, routines...) {
		if i == false {
			distance++
		}
	}
	return distance, nil
}
