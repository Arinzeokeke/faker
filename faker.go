package faker

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"../../tr"
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
	locale        string
	defaultLocale string
	engine        *tr.Engine
}

// Faker interface
type Faker interface {
	Name() Name
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
func (f *Fake) Name() Name {
	return Name{
		f:      f,
		engine: f.engine,
	}
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
