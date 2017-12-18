package faker

// Namer interface
type Namer interface {
	FirstName() string
	LastName() string
	FullName() string
	JobTitle() string
	Prefix() string
	Suffix() string
	JobDescriptor() string
	JobArea() string
	JobType() string
	Gender() string
}

// Name struct
type Name struct {
	*Fake
}

// FirstName returns first name
func (n *Name) FirstName() string {
	if directoryExists(n.DefaultLocale+"/"+namePrefix+"/male_first_name") &&
		directoryExists(n.DefaultLocale+"/"+namePrefix+"/female_first_name") {

		// TODO give gender choice
		if !directoryExists(n.DefaultLocale + "/" + namePrefix + "/first_name") {
			return n.pick(namePrefix + "/female_first_name")
		}

	}
	return n.pick(namePrefix + "/first_name")
}

// LastName returns last name
func (n *Name) LastName() string {
	if directoryExists(n.DefaultLocale+"/"+namePrefix+"/male_last_name") &&
		directoryExists(n.DefaultLocale+"/"+namePrefix+"/female_last_name") {

		// TODO give gender choice
		if !directoryExists(n.DefaultLocale + "/" + namePrefix + "/last_name") {
			return n.pick(namePrefix + "/female_last_name")
		}

	}
	return n.pick(namePrefix + "/last_name")
}

// FullName gets full name
func (n *Name) FullName() string {
	return n.FirstName() + " " + n.LastName()
}

// JobDescriptor returns a job description
func (n *Name) JobDescriptor() string {
	return n.pick(namePrefix + "/title/descriptor")
}

//JobArea returns a job area
func (n *Name) JobArea() string {
	return n.pick(namePrefix + "/title/level")
}

//JobType returns a job type
func (n *Name) JobType() string {
	return n.pick(namePrefix + "/title/job")
}

//JobTitle returns a job title
func (n *Name) JobTitle() string {
	return n.JobDescriptor() + " " + n.JobArea() + " " + n.JobType()
}

// Gender returns a gender
func (n *Name) Gender() string {
	return n.pick(namePrefix + "/gender")
}

//Prefix returns a prefix
func (n *Name) Prefix() string {
	// TODO handle gender
	return n.pick(namePrefix + "/prefix")
}

//Suffix returns a suffix
func (n *Name) Suffix() string {
	// TODO handle gender
	return n.pick(namePrefix + "/suffix")
}
