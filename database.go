package faker

// Databaser Interface
type Databaser interface {
	Column() string
	Type() string
	Collation() string
	Engine() string
}

// Database struct
type Database struct {
	*Fake
}

// Column returns a database column
func (d *Database) Column() string {
	return d.pick(databasePrefix + "/column")
}

// Type returns a database type
func (d *Database) Type() string {
	return d.pick(databasePrefix + "/type")
}

// Collation returns a database collation
func (d *Database) Collation() string {
	return d.pick(databasePrefix + "/collation")
}

// Engine returns a database Engine
func (d *Database) Engine() string {
	return d.pick(databasePrefix + "/engine")
}
