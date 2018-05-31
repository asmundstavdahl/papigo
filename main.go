package main

import (
	"flag"
	"fmt"
	"time"
)

var pi float64
var div float64
var iteration uint

var flagOutputInterval uint
var flagBufferLength uint
var flagParallellism uint
var flagChunkSize uint

func main() {
	pi = 0.0
	div = 1.0

	flag.UintVar(&flagOutputInterval, "output-interval", 5, "seconds between each printing of progress")
	flag.UintVar(&flagBufferLength, "buffer", 10, "size of channel buffers")
	flag.UintVar(&flagParallellism, "parallellism", 10, "number of goroutine pairs to instantiate")
	flag.UintVar(&flagChunkSize, "chunk", 10000, "calculations before sending in channel")

	flag.Parse()

	chAdd := make(chan float64, flagBufferLength)
	chSub := make(chan float64, flagBufferLength)

	if flagParallellism > 0 {
		skips := 4.0 * float64(flagParallellism)
		for i := 0; i < int(flagParallellism); i++ {
			go divider(chAdd, div+(float64(i*2)*2.0), skips)
			go divider(chSub, div+(float64(i*2+1)*2.0), skips)
		}

		go sum(chAdd, chSub)
	} else {
		go singleThread()
	}

	fmt.Printf("%s\t%s\n", "iter.", "pi")

	if flagOutputInterval > 0 {
		for {
			<-time.After(time.Duration(flagOutputInterval) * time.Second)
			fmt.Printf("%d\t%.54f\n", iteration, pi)
		}
	} else {
		for {
			<-time.After(time.Second)
		}
	}
}

func divider(ch chan float64, div float64, skips float64) {
	var tmp float64
	for {
		tmp = 0
		for i := uint(0); i < flagChunkSize; i++ {
			tmp += 4.0 / div
			div += float64(skips)
		}
		ch <- tmp
	}
}

func sum(chAdd chan float64, chSub chan float64) {
	for {
		pi += <-chAdd
		pi -= <-chSub
		iteration += flagChunkSize
	}
}

func singleThread() {
	for {
		for i := uint(0); i < flagChunkSize; i++ {
			pi += 4.0 / div
			div += 2.0
			pi -= 4.0 / div
			div += 2.0
		}
		iteration += flagChunkSize
		<-time.After(1 * time.Nanosecond)
	}
}
