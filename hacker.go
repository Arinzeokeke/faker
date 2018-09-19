package faker

// Hackier Interface
type Hackier interface {
	Abbreviation() string
	Adjective() string
	Noun() string
	Verb() string
	Ingverb() string
	Phrase() string
}

// Hacker struct
type Hacker struct {
	*Fake
}

// Abbreviation Returns an abbreviation
func (h *Hacker) Abbreviation() string {
	return h.pick(hackerPrefix + "/abbreviation")
}

// Adjective Returns an adjective
func (h *Hacker) Adjective() string {
	return h.pick(hackerPrefix + "/adjective")
}

// Noun Returns a noun
func (h *Hacker) Noun() string {
	return h.pick(hackerPrefix + "/noun")
}

// Verb Returns a verb
func (h *Hacker) Verb() string {
	return h.pick(hackerPrefix + "/verb")
}

// Ingverb Returns an ingverb
func (h *Hacker) Ingverb() string {
	return h.pick(hackerPrefix + "/ingverb")
}

// Phrase Returns a phrase made up of other hacker verbs
func (h *Hacker) Phrase() string {
	hackerData := map[string]string{
		"abbreviation": h.Hacker().Abbreviation(),
		"adjective":    h.Hacker().Adjective(),
		"ingverb":      h.Hacker().Ingverb(),
		"noun":         h.Hacker().Noun(),
		"verb":         h.Hacker().Verb(),
	}
	phrase := h.pick(hackerPrefix + "/phrase")

	return mustache(phrase, hackerData)
}
