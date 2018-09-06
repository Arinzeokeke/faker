package faker

// Internetier interface
type Internetier interface {
	Avatar() string
	Protocol() string
	DomainSuffix() string
	Email() string
	ExampleEmail() string
}

// Internet struct
type Internet struct {
	*Fake
}

// Avatar Generates a URL for an avatar.
func (i *Internet) Avatar() string {
	return i.pick(internetPrefix + "/avatar_uri")
}
