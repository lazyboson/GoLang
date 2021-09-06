package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	//output turns out to be zero, so line 11, 12 are being executed prior to line 9
	//“Most of the time, data races are introduced because the developers are thinking about the problem sequentially.
	//They assume that because a line of code falls before another that it will run first”

	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	} else {
		fmt.Printf("the value is %v.\n", data)
	}

	//oneway to solve is using the resourse synchronization
	//Mutex is only making sure -- that only the one would be executed after other--but mutex is not solving which should be executed first
	//so race condition is still remained unsolved
	var memoryAccess sync.Mutex
	var data1 int
	go func() {
		memoryAccess.Lock()
		data1++
		memoryAccess.Unlock()
	}()
	memoryAccess.Lock()
	if data1 == 0 {
		fmt.Printf("the value is %v.\n", data1)
	} else {
		fmt.Printf("the value is %v.\n", data1)
	}
	memoryAccess.Unlock()

	//deadlock -- “A deadlocked program is one in which all concurrent processes are waiting on one another”

	type value struct {
		mu    sync.Mutex
		value int
	}
	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()
		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()

	//coffman Conditions for the deadlock -- help, prevent and correct deadlock
	//Livelock --“Livelocks are programs that are actively performing concurrent operations,
	//but these operations do nothing to move the state of the program forward”

	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(1 * time.Millisecond) {
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}
	tryDir := func(DirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprint(out, " %v", DirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, " . Success!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryleft := func(out *bytes.Buffer) bool {
		return tryDir("left", &left, out)
	}
	tryright := func(out *bytes.Buffer) bool {
		return tryDir("right", &right, out)
	}
	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() {
			fmt.Println(out.String())
		}()
		defer walking.Done()
		fmt.Fprintf(&out, " %v is trying to scoot:", name)
		for i := 0; i < 5; i++ {
			if tryleft(&out) || tryright(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallWay sync.WaitGroup
	peopleInHallWay.Add(2)
	go walk(&peopleInHallWay, "Alice")
	go walk(&peopleInHallWay, "Bob")
	peopleInHallWay.Wait()
}
