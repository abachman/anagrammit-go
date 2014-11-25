// It's possible to run the generator in multiple goroutines simultaneously

package main

import (
	"flag"
	"fmt"
	"math/rand"

	"github.com/abachman/anagrammit-go/anagrammit"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	threads := flag.Int("threads", 3, "goroutine count")
	flag.Parse()

	args := &anagrammit.GeneratorArgs{
		WordLength:  3,
		ResultLimit: 5,
		Shuffle:     true,
		WordsFile:   "tmp/dictionary.txt",
	}

	// Base
	generator := anagrammit.NewGenerator(args)

	quit := make(chan bool, *threads)

	for i := 0; i < *threads; i++ {
		// first
		go func(n int) {
			output := make(chan string)
			generator.Generate("hello world "+randSeq(6), output)

			for msg := range output {
				fmt.Printf("[%v] %s\n", n, msg)
			}

			quit <- true
		}(i)
	}

	for i := 0; i < *threads; i++ {
		<-quit
	}
}
