package faker

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"../../tr"
)

const (
	namePrefix = "name"
	addressPrefix = "address"
)

var (
	definitions = map[string][]string{
		"name":         []string{"first_name", "last_name", "prefix", "suffix", "gender", "title", "male_first_name", "female_first_name", "male_middle_name", "female_middle_name", "male_last_name", "female_last_name"},
		"address":      []string{"city_prefix", "city_suffix", "street_suffix", "county", "country", "country_code", "state", "state_abbr", "street_prefix", "postcode", "direction", "direction_abbr"},
		"company":      []string{"adjective", "noun", "descriptor", "bs_adjective", "bs_noun", "bs_verb", "suffix"},
		"lorem":        []string{"words"},
		"hacker":       []string{"abbreviation", "adjective", "noun", "verb", "ingverb", "phrase"},
		"phone_number": []string{"formats"},
		"finance":      []string{"account_type", "transaction_type", "currency", "iban", "credit_card"},
		"internet":     []string{"avatar_uri", "domain_suffix", "free_email", "example_email", "password"},
		"commerce":     []string{"color", "department", "product_name", "price", "categories"},
		"database":     []string{"collation", "column", "engine", "type"},
		"system":       []string{"mimeTypes"},
		"date":         []string{"month", "weekday"},
		"title":        []string{},
		"separator":    []string{},
	}
)

// Fake Object
type Fake struct {
	Locale        string
	DefaultLocale string
	Engine        *tr.Engine
}

// Faker interface
type Faker interface {
	Name() *Name
	Lang(l string)
}

//New creates new faker
func New() (*Fake, error) {

	e, err := tr.Init("../locale", "en")
	if err != nil {
		return nil, fmt.Errorf("Couldn't load locale, Error: %v", err)

	}

	fake := &Fake{
		Locale:        "en",
		DefaultLocale: "en",
		Engine:        e,
	}

	return fake, nil

}

// Name returns name
func (f *Fake) Name() *Name {
	return &Name{f}
}

//Lang sets lang
func (f *Fake) Lang(l string) *Fake {
	f.Engine.DefaultLocale = f.Engine.Lang(l)
	f.Locale = l
	return f
}

func random(i int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(i)
}

func directoryExists(dir string) bool {
	dir, _ = filepath.Abs(dir)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}
