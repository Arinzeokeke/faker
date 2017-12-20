package faker

import (
	"fmt"
)

// Companier Interface
type Companier interface {
	CompanyName() string
	CompanySuffix() string
	CatchPhrase() string
	Bs() string
	CatchPhraseAdjective() string
	CatchPhraseDescriptor() string
	CatchPhraseNoun() string
	BsAdjective() string
	BsBuzz() string
	BsNoun() string
}

// Company struct
type Company struct {
	*Fake
}

// CompanyName returns a company name
func (c *Company) CompanyName() string {
	return c.format(c.pick(companyPrefix+"/name"), companyPrefix)
}

// CompanySuffix returns a company suffix
func (c *Company) CompanySuffix() string {
	return c.pick(companyPrefix + "/suffix")
}

// CatchPhrase returns a catch phrase
func (c *Company) CatchPhrase() string {
	return fmt.Sprintf("%s %s %s", c.CatchPhraseAdjective(), c.CatchPhraseDescriptor(), c.CatchPhraseNoun())
}

// Bs returns a bs
func (c *Company) Bs() string {
	return fmt.Sprintf("%s %s %s", c.BsBuzz(), c.BsAdjective(), c.BsNoun())
}

// CatchPhraseAdjective returns a company catchphrase adjective
func (c *Company) CatchPhraseAdjective() string {
	return c.pick(companyPrefix + "/adjective")
}

// CatchPhraseDescriptor returns a company catchphrase descriptor
func (c *Company) CatchPhraseDescriptor() string {
	return c.pick(companyPrefix + "/descriptor")
}

// CatchPhraseNoun returns a company catchphrase noun
func (c *Company) CatchPhraseNoun() string {
	return c.pick(companyPrefix + "/noun")
}

// BsAdjective returns a bs adjective
func (c *Company) BsAdjective() string {
	return c.pick(companyPrefix + "/bs_adjective")
}

// BsBuzz returns a bs buzz
func (c *Company) BsBuzz() string {
	return c.pick(companyPrefix + "/bs_verb")
}

// BsNoun returns a bs Noun
func (c *Company) BsNoun() string {
	return c.pick(companyPrefix + "/bs_noun")
}
