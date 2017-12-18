package faker

import "github.com/Arinzeokeke/tr"
import "strings"

const (
	prefix = "name"
)

// Namer interface
type Namer interface {
	FirstName() string
	// LastName() string
	// FindName() string
	// JobTitle() string
	// Prefix() string
	// Suffix() string
	// Title() string
	// JobDescriptor() string
	// JobArea() string
	// JobType() string
}

// Name struct
type Name struct {
	engine *tr.Engine
	f      Faker
}

// FirstName returns first name
func (n *Name) FirstName() string {

	firsties := strings.Split(n.engine.Tr(prefix+"/first_name"), "\n")
	return firsties[0]
}
