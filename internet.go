package faker

import "strings"

// Internetier interface
type Internetier interface {
	Avatar() string
	DomainSuffix() string
	Email() string
	ExampleEmail() string
	Protocol() string
	UserName() string
}

// Internet struct
type Internet struct {
	*Fake
}

// Avatar Generates a URL for an avatar.
func (i *Internet) Avatar() string {
	return i.pick(internetPrefix + "/avatar_uri")
}

// DomainSuffix Generates a random domain suffix.
func (i *Internet) DomainSuffix() string {
	return i.pick(internetPrefix + "/domain_suffix")
}

// Protocol Randomly generates http or https
func (i *Internet) Protocol() string {
	protocols := []string{"http", "https"}
	return selectElement(protocols)
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
