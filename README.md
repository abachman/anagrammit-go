# anagrammit-go -- Anagram Generation in Go

A pure Go anagram generation library. Useful for building command line tools or
web services. Bring your own [word lists](http://www.anagrammy.com/resources/wordlists.html).

An example.

```go
// A basic command line anagram generator.

package main

import (
	"flag"
	"fmt"

	"github.com/abachman/anagrammit-go/anagrammit"
)

var p = flag.Parse

func main() {
	wordLen := flag.Int("wordlength", 3, "minimum word length")
	limit := flag.Int("limit", 10, "result limit, use 0 for unlimited")
	shuffle := flag.Bool("shuffle", false, "shuffle lexicon")
	flag.Parse()
	inpt := flag.Arg(0)

	args := &anagrammit.GeneratorArgs{
		WordLength:  *wordLen,
		ResultLimit: *limit,
		Shuffle:     *shuffle,
		WordsFile:   "tmp/dictionary.txt",
	}

	// Base
	generator := anagrammit.NewGenerator(args)

	output := make(chan string)
	generator.Generate(inpt, output)

	for msg := range output {
		fmt.Println(msg)
	}
}
```

Or, download and build the examples:

    $ make generator
    $ ./anagram-generator -limit=0 "pure soap union"
