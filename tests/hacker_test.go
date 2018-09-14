package faker

import (
	"strings"
	"testing"
)

func TestAbbreviation(t *testing.T) {
	abb := f.Hacker().Abbreviation()

	if abb != "" {
		return
	}

	t.Error("Could not generate hacker abbreviation")
}

func TestAdjective(t *testing.T) {
	adj := f.Hacker().Adjective()

	if adj != "" {
		return
	}

	t.Error("Could not generate hacker adjective")
}

func TestNoun(t *testing.T) {
	noun := f.Hacker().Noun()

	if noun != "" {
		return
	}

	t.Error("Could not generate hacker noun")
}

func TestVerb(t *testing.T) {
	verb := f.Hacker().Verb()

	if verb != "" {
		return
	}

	t.Error("Could not generate hacker verb")
}

func TestIngverb(t *testing.T) {
	ingverb := f.Hacker().Ingverb()

	if ingverb != "" {
		return
	}

	t.Error("Could not generate hacker -ingverb")
}

func TestPhrase(t *testing.T) {
	phrase := f.Hacker().Phrase()
	expectedLen := 5

	splitPhrase := strings.Split(phrase, " ")
	if len(splitPhrase) >= expectedLen {
		return
	}

	t.Errorf("Expected hacker phrase of minimum length %v, got %v", expectedLen, len(splitPhrase))
}
