package anagrammit

type Word struct {
	Display     string
	LetterCount []int
}

//// Word ///////////////////

func NewWord(inWord string) *Word {
	return &Word{inWord, letterFrequency(inWord)}
}
