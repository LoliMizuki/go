package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	close_channel()
	// select_channel()
}

func close_channel() {
	chVars := make(chan int, 10)

	for i := 0; i < 10; i++ {
		chVars <- i
	}
	close(chVars)

	for cv := range chVars {
		fmt.Println(cv)
	}
}

func select_channel() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	quit := make(chan int)

	go random_set_ch_value_ten_times(c1, c2, c3, quit)

	for {
		select {
		case <-c1:
			fmt.Println("I like c1")
		case i := <-c2:
			fmt.Println("I hate c2", i)
		case <-c3:
			fmt.Println("Fear is CUTE !!!! c3")
		case <-quit:
			fmt.Println("game over")
			return
		}
	}
}

func random_set_ch_value_ten_times(c1, c2, c3, quit chan int) {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < 10; i++ {
		switch choice := rand.Int31n(3); choice {
		case 0:
			c1 <- 1
		case 1:
			c2 <- 1
		case 2:
			c3 <- 1
		}

		time.Sleep(2 * time.Nanosecond)
	}

	quit <- 1
}
