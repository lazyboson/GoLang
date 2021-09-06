package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine is sleeping...")
		time.Sleep(1)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine is sleeping....")
		time.Sleep(1)
	}()
	//stopping main thread from exiting prior completing all the go routines
	wg.Wait()
	fmt.Println("All goroutines are completed")

	var count int //default == 0
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)

	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)

	}

	//increment
	var airthmatic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		airthmatic.Add(1)
		go func() {
			defer airthmatic.Done()
			increment()
		}()
	}

	//decrement
	for i := 0; i <= 5; i++ {
		airthmatic.Add(1)
		go func() {
			defer airthmatic.Done()
			decrement()
		}()
	}

	airthmatic.Wait()
	fmt.Println("Airthmatic completed...")

	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		for i := 5; i > 0; i-- {
			l.Lock()
			l.Unlock()
			time.Sleep(1)
		}

	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()
	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(tw, "%d\t%v\t%v\n", count, test(count, &m, m.RLocker()), test(count, &m, &m))
	}

}
