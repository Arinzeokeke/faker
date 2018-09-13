package faker

import "strings"

// Loremer Interface
type Loremer interface {
	Sentence(wordCount int) string
	Sentences(sentenceCount int, separator string) string
	Word() string
	Words(num int) string
}

// Lorem struct
type Lorem struct {
	*Fake
}

// Sentence Generates a sentence of n words
func (l *Lorem) Sentence(wordCount int) string {
	sentence := l.Lorem().Words(wordCount)
	return strings.ToUpper(string(sentence[0])) + string(sentence[1:]) + "."
}

// Sentences Generate sentences divided by a separator
func (l *Lorem) Sentences(sentenceCount int, separator string) string {
	sentences := []string{}
	for count := sentenceCount; count > 0; count-- {
		wc := randomIntRange(3, 10)
		sentences = append(sentences, l.Lorem().Sentence(wc))
	}
	return strings.Join(sentences, separator)
}

// Word Generates a random word
func (l *Lorem) Word() string {
	return l.pick(loremPrefix + "/words")
}

// Words Generates an amount of words
func (l *Lorem) Words(num int) string {
	words := []string{}
	for i := 0; i < num; i++ {
		words = append(words, l.Lorem().Word())
	}
	return strings.Join(words, " ")
}
