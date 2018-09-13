package faker

import "strings"

// Loremer Interface
type Loremer interface {
	Lines(lineCount int) string
	Paragraph(sentenceCount int) string
	Paragraphs(sentenceCount int, separator string) string
	Sentence(wordCount int) string
	Sentences(sentenceCount int, separator string) string
	Slug(wordCount int) string
	Word() string
	Words(num int) string
}

// Lorem struct
type Lorem struct {
	*Fake
}

// Lines Generates lines of lorem separated by `'\n'`
func (l *Lorem) Lines(lineCount int) string {
	return l.Lorem().Sentences(lineCount, "\n")
}

// Paragraph Generate a paragraph
func (l *Lorem) Paragraph(sentenceCount int) string {
	return l.Lorem().Sentences(sentenceCount+random(3), " ")
}

// Paragraphs Generate n paragraph(s)
func (l *Lorem) Paragraphs(sentenceCount int, separator string) string {
	paragraphs := []string{}
	for count := sentenceCount; count > 0; count-- {
		paragraphs = append(paragraphs, l.Lorem().Paragraph(3))
	}
	return strings.Join(paragraphs, "\n \r")
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

// Slug Generate a slugged sentence
func (l *Lorem) Slug(wordCount int) string {
	return slugify(l.Lorem().Words(wordCount))
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
