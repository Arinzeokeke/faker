package faker

type Namer interface {
	FirstName() string
	LastName() string
	FindName() string
	JobTitle() string
	Prefix() string
	Suffix() string
	Title() string
	JobDescriptor() string
	JobArea() string
	JobType() string
}

type Name struct{}
