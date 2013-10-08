package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// wait_goroutine_completed()
	// wait_multi_goroutines_completed()
	time_out()
}

func wait_goroutine_completed() {
	chRoutineDone := make(chan int)
	go func() {
		time.Sleep(5 * 1e9)
		fmt.Println("I sleep for the Horde")
		chRoutineDone <- 1
	}()

	<-chRoutineDone

	fmt.Println("done")
}

func wait_multi_goroutines_completed() {
	numberOfChannels := 10

	chDones := make(chan int)

	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)

	for i := 0; i < numberOfChannels; i++ {
		go func(index int) {
			sleepSec := rand.Float64() * 3
			time.Sleep(time.Duration(sleepSec * 1e9))

			fmt.Printf("#%d is done(%.2f)\n", index, sleepSec)
			chDones <- 1
		}(i)
	}

	for i := 0; i < numberOfChannels; i++ {
		<-chDones
	}
}

func time_out() {
	throughput := make(chan int)
	timeOut := make(chan bool, 1)

	rand.Seed(int64(time.Now().Nanosecond()))

	go func() {
		timeSec := rand.Float64() * 2
		time.Sleep(time.Duration(timeSec * 1e9))

		if timeSec >= 1.0 {
			timeOut <- true
		} else {
			// throughput <- rand.Intn(10)
			throughput <- 100
		}
	}()

	var iValue int
	select {
	case iValue = <-throughput:
		fmt.Println("got result: ", iValue)
	case <-timeOut:
		fmt.Println("time out")
	}
}
