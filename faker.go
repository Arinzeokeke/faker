package faker

import (
	"fmt"

	"github.com/Arinzeokeke/tr"
)

// Fake Object
type Fake struct {
	locale        string
	defaultLocale string
	engine        *tr.Engine
}

// Faker interface
type Faker interface {
	Name() Namer
}

//New creates new faker
func New() (*Fake, error) {

	engine, err := tr.Init("locale", "en")
	if err != nil {
		return nil, fmt.Errorf("Couldn't load locale, Error: %v", err)

	}

	fake := &Fake{
		locale:        "en",
		defaultLocale: "en",
		engine:        engine,
	}

	return fake, nil

}

// Name returns name
func (f *Fake) Name() Namer {
	return Name{
		f:      f,
		engine: f.engine,
	}

}
