package faker

// Addresser interface
type Addresser interface {
	ZipCode() string
	// City() string
	// CityPrefix() string
	// CitySuffix() string
	// StreetName() string
	// StreetAddress() string
	// StreetSuffix() string
	// StreetPrefix() string
	// SecondaryAddress() string
	// County() string
	// Country() string
	// CountryCode() string
	// State() string
	// StateAbbr() string
	// Latitude() float64
	// Longitude() float64
}

// Address struct
type Address struct {
	*Fake
}

//ZipCode returns a zip code
func (a *Address) ZipCode() string {
	// TODO handle gender
	return a.pick("/suffix")
}
