// A basic anagram generating web service

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/abachman/anagrammit-go/anagrammit"
)

var generator *anagrammit.Generator

func anagramHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	phrase := r.FormValue("phrase")

	log.Println("responding to input phrase", phrase)

	if len(phrase) > 0 {
		output := make(chan string)
		generator.Generate(phrase, output)
		for msg := range output {
			fmt.Fprintln(w, msg)
		}
	} else {
		fmt.Fprintln(w, "")
	}
}

func main() {
	wordLen := flag.Int("wordlength", 3, "minimum word length")
	limit := flag.Int("limit", 10, "result limit")
	shuffle := flag.Bool("shuffle", false, "shuffle lexicon")
	flag.Parse()

	args := &anagrammit.GeneratorArgs{
		WordLength:  *wordLen,
		ResultLimit: *limit,
		Shuffle:     *shuffle,
		WordsFile:   "tmp/common-word-list.txt",
	}

	generator = anagrammit.NewGenerator(args)

	http.HandleFunc("/generate", anagramHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
