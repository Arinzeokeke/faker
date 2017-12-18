package faker

import "../../tr"
import "strings"

const (
	prefix = "name"
)

// Namer interface
type Namer interface {
	FirstName() string
	LastName() string
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
	f      *Fake
}

// FirstName returns first name
func (n *Name) FirstName() string {
	if directoryExists(n.f.defaultLocale+"/"+prefix+"/male_first_name") &&
		directoryExists(n.f.defaultLocale+"/"+prefix+"/female_first_name") {

		// TODO give gender choice
		if !directoryExists(n.f.defaultLocale + "/" + prefix + "/first_name") {
			firsties, err := getList(n.engine, prefix+"/female_first_name", n.f.defaultLocale)
			if err != nil {
				panic(err)
			}
			return firsties[random(len(firsties))]
		}

	}
	firsties, err := getList(n.engine, prefix+"/first_name", n.f.defaultLocale)
	if err != nil {
		panic(err)
	}
	return firsties[random(len(firsties))]
}

// LastName returns last name
func (n *Name) LastName() string {
	if directoryExists(n.f.defaultLocale+"/"+prefix+"/male_last_name") &&
		directoryExists(n.f.defaultLocale+"/"+prefix+"/female_last_name") {

		// TODO give gender choice
		if !directoryExists(n.f.defaultLocale + "/" + prefix + "/last_name") {
			firsties, err := getList(n.engine, prefix+"/female_last_name", n.f.defaultLocale)
			if err != nil {
				panic(err)
			}
			return firsties[random(len(firsties))]
		}

	}
	lasties, err := getList(n.engine, prefix+"/last_name", n.f.defaultLocale)
	if err != nil {
		panic(err)
	}
	return lasties[random(len(lasties))]
}

func getList(n *tr.Engine, q string, d string) ([]string, error) {
	w, err := n.Tr(q)
	if err != nil {
		w, err = n.Lang(d).Tr(q)
	}
	if err != nil {
		return nil, err
	}
	return strings.Split(w, "\n"), nil
}
