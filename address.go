package faker

import (
	"strings"
)

// Addresser interface
type Addresser interface {
	ZipCode() string
	City() string
	CityPrefix() string
	CitySuffix() string
	StreetName() string
	StreetAddress() string
	StreetSuffix() string
	StreetPrefix() string
	SecondaryAddress() string
	County() string
	Country() string
	CountryCode() string
	State() string
	StateAbbr() string
	Latitude() float64
	Longitude() float64
	Direction(bool) string
	OrdinalDirection(bool) string
	CardinalDirection(bool) string
}

// Address struct
type Address struct {
	*Fake
}

//ZipCode returns a zip code
func (a *Address) ZipCode() string {
	// TODO handle format choices
	return replaceSymbols(a.pick(addressPrefix + "/postcode"))
}

//City returns a city
func (a *Address) City() string {
	return a.format(a.pick(addressPrefix+"/city"), addressPrefix)

}

//CityPrefix returns a city prefix
func (a *Address) CityPrefix() string {
	return a.pick(addressPrefix + "/city_prefix")
}

//CitySuffix returns a city suffix
func (a *Address) CitySuffix() string {
	return a.pick(addressPrefix + "/city_suffix")
}

//StreetName returns a street name
func (a *Address) StreetName() string {
	s := " " + a.StreetSuffix()

	switch random(2) {
	case 0:
		return a.Name().LastName() + s
	case 1:
		return a.Name().FirstName() + s
	}
	return a.Name().FirstName() + s

}

//StreetAddress returns a street address
func (a *Address) StreetAddress() string {
	return strings.Repeat("#", random(6))

}

//StreetPrefix returns a street prefix
func (a *Address) StreetPrefix() string {
	return a.pick(addressPrefix + "/street_prefix")
}

//StreetSuffix returns a street suffix
func (a *Address) StreetSuffix() string {
	return a.pick(addressPrefix + "/street_suffix")
}

//SecondaryAddress returns a secondary address
func (a *Address) SecondaryAddress() string {
	opts := []string{"Apt. ###", "Suite ###"}
	return replaceSymbols(opts[random(len(opts))])
}

//County returns a county
func (a *Address) County() string {
	return a.pick(addressPrefix + "/county")
}

//Country returns a country
func (a *Address) Country() string {
	return a.pick(addressPrefix + "/country")
}

//CountryCode returns a country code
func (a *Address) CountryCode() string {
	return a.pick(addressPrefix + "/country_code")
}

//State returns a state
func (a *Address) State() string {
	return a.pick(addressPrefix + "/state")
}

//StateAbbr returns a state abbreviation
func (a *Address) StateAbbr() string {
	return a.pick(addressPrefix + "/state_abbr")
}

//Latitude returns a latitude
func (a *Address) Latitude() float64 {
	//TODO
	return 4.555
}

//Longitude returns a longitude
func (a *Address) Longitude() float64 {
	//TODO
	return 4.555
}

//Direction returns a navigational direction
func (a *Address) Direction(abbr bool) string {
	if abbr {
		return a.pick(addressPrefix + "/direction_abbr")
	}
	return a.pick(addressPrefix + "/direction")
}

//CardinalDirection returns a navigational cardinal direction
func (a *Address) CardinalDirection(abbr bool) string {
	//TODO handle cardinal  0, 4
	if abbr {
		return a.pick(addressPrefix + "/direction_abbr")
	}
	return a.pick(addressPrefix + "/direction")
}

//OrdinalDirection returns a navigational ordinal direction
func (a *Address) OrdinalDirection(abbr bool) string {
	//TODO handle ordinal 4, 8
	if abbr {
		return a.pick(addressPrefix + "/direction_abbr")
	}
	return a.pick(addressPrefix + "/direction")
}
