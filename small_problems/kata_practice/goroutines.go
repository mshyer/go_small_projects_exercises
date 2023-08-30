package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

var counter = 0


func main() {
	for i := 0; i < 20; i++ {
		go printSomeVal()
	}

	time.Sleep(time.Millisecond * 10)
}

func printSomeVal() {
	randVal := rand.Int31n(1000)
	fmt.Println(randVal)
}