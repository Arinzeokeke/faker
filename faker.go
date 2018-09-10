package faker

import (
	"fmt"
	"strings"

	"github.com/Arinzeokeke/tr"
)

const (
	namePrefix     = "name"
	addressPrefix  = "address"
	commercePrefix = "commerce"
	companyPrefix  = "company"
	databasePrefix = "database"
	datePrefix     = "date"
	internetPrefix = "internet"
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
	Name() Namer
	Lang(l string)
	format(s, p string) string
	Address() Addresser
	Commerce() Commercer
	Company() Companier
	Database() Databaser
	Dater() Dater
	Internet() Internetier
}

// Config for Faker
type Config struct {
	defaultLocale string
}

//New creates new faker
func New(c *Config) (*Fake, error) {

	e, err := tr.Init("locale", c.defaultLocale)
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
func (f *Fake) Name() Namer {
	return &Name{f}
}

// Address returns address
func (f *Fake) Address() Addresser {
	return &Address{f}
}

// Commerce returns commerce
func (f *Fake) Commerce() Commercer {
	return &Commerce{f}
}

// Company returns company
func (f *Fake) Company() Companier {
	return &Company{f}
}

// Database returns database
func (f *Fake) Database() Databaser {
	return &Database{f}
}

// Date returns date
func (f *Fake) Date() Dater {
	return &Date{f}
}

// Internet returns internet
func (f *Fake) Internet() Internetier {
	return &Internet{f}
}

//Lang sets lang
func (f *Fake) Lang(l string) *Fake {
	f.Engine.DefaultLocale = f.Engine.Lang(l)
	f.Locale = l
	return f
}

func (f *Fake) format(s, p string) string {
	for strings.Index(s, "#{") != -1 && strings.Index(s, "}") != 1 {
		start := strings.Index(s, "#{")
		end := strings.Index(s, "}")
		token := s[start+2 : end-1]
		token = strings.ToLower(token)
		token = strings.Replace(s, ".", "/", -1)
		var affix string
		if strings.Contains(token, "/") {
			affix = token
		} else {
			affix = p + "/" + token
		}
		out := f.pick(affix)
		s = s[:start] + out + s[end+1:]
	}
	return s
}

func (f *Fake) pick(affix string) string {
	v, err := getList(f.Engine, affix, f.DefaultLocale)
	if err != nil {
		panic(err)
	}
	return v[random(len(v))]

}
