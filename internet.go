package faker

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

// Internetier interface
type Internetier interface {
	Avatar() string
	Color(baseRed255, baseGreen255, baseBlue255 float64) string
	DomainName() string
	DomainSuffix() string
	DomainWord() string
	Email(firstName, lastName, provider string) string
	ExampleEmail(firstName, lastName string) string
	IP() string
	IPv6() string
	Mac(sep string) string
	Protocol() string
	URL() string
	UserName(firstName, lastName string) string
}

// Internet struct
type Internet struct {
	*Fake
}

// Avatar Generates a URL for an avatar.
func (i *Internet) Avatar() string {
	return i.pick(internetPrefix + "/avatar_uri")
}

// Color Generates a random hexadecimal color.
// based on awesome response : http://stackoverflow.com/questions/43044/algorithm-to-randomly-generate-an-aesthetically-pleasing-color-palette
func (i *Internet) Color(baseRed255, baseGreen255, baseBlue255 float64) string {
	red := math.Floor((float64(random(256)) + baseRed255) / 2)
	green := math.Floor((float64(random(256)) + baseRed255) / 2)
	blue := math.Floor((float64(random(256)) + baseRed255) / 2)

	redStr := fmt.Sprintf("%02X", int(red))
	greenStr := fmt.Sprintf("%02X", int(green))
	blueStr := fmt.Sprintf("%02X", int(blue))

	return "#" + redStr + greenStr + blueStr
}

// DomainName Generates a random domain name.
func (i *Internet) DomainName() string {
	return i.Internet().DomainWord() + "." + i.Internet().DomainSuffix()
}

// DomainSuffix Generates a random domain suffix.
func (i *Internet) DomainSuffix() string {
	return i.pick(internetPrefix + "/domain_suffix")
}

// DomainWord Generates a random domain word.
func (i *Internet) DomainWord() string {
	re := regexp.MustCompile(`([\\~#&*{}/:<>?|\"'])`)
	return strings.ToLower(re.ReplaceAllString(i.Name().FirstName(), ""))
}

// Email Generates a valid email address
func (i *Internet) Email(firstName, lastName, provider string) string {
	if provider == "" {
		provider = i.pick(internetPrefix + "/free_email")
	}
	return slugify(strings.ToLower(i.Internet().UserName(firstName, lastName))) + "@" + provider
}

// ExampleEmail Provides an email example
func (i *Internet) ExampleEmail(firstName, lastName string) string {
	provider := i.pick(internetPrefix + "/example_email")
	return i.Internet().Email(firstName, lastName, provider)
}

// IP Generates a random IP
func (i *Internet) IP() string {
	result := []string{}

	for i := 0; i < 4; i++ {
		result = append(result, string(random(255)))
	}

	return strings.Join(result, ".")
}

// IPv6 Generates a random IPv6 address
func (i *Internet) IPv6() string {
	result := []string{}

	for i := 0; i < 8; i++ {
		result = append(result, randHash())
	}

	return strings.Join(result, ".")
}

// Mac Generates a random mac address.
func (i *Internet) Mac(sep string) string {
	mac := ""
	validSep := ":"

	// if the client passed in a different separator than `:`,
	// we will use it if it is in the list of acceptable separators (dash or no separator)
	acceptedSep := strslice{"-", ""}
	if acceptedSep.indexOf(sep) != -1 {
		validSep = sep
	}

	for i := 0; i < 12; i++ {
		mac += strings.TrimSpace(fmt.Sprintf("%2x", random(15)))

		if (i%2 == 1) && (i != 11) {
			mac += validSep
		}
	}

	return mac
}

// Protocol Randomly generates http or https
func (i *Internet) Protocol() string {
	protocols := []string{"http", "https"}
	return selectElement(protocols)
}

// URL Generates a random url.
func (i *Internet) URL() string {
	return i.Internet().Protocol() + "://" + i.Internet().DomainName()
}

// UserName Generates a username based on one of several patterns
func (i *Internet) UserName(firstName, lastName string) string {
	var result string

	if firstName == "" {
		firstName = i.Name().FirstName()
	}

	if lastName == "" {
		lastName = i.Name().LastName()
	}

	switch random(2) {
	case 0:
		result = firstName + string(random(99))
	case 1:
		result = firstName + selectElement([]string{".", "_"}) + lastName
	case 2:
		result = firstName + selectElement([]string{".", "_"}) + lastName + string(random(99))
	}

	result = strings.Replace(result, "'", "", -1)
	result = strings.Replace(result, " ", "", -1)

	return result
}
