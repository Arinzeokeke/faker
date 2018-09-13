package faker

import (
	"regexp"
	"testing"
)

func init() {
	if err != nil {
		panic(err)
	}
}

func TestCollation(t *testing.T) {
	collation := f.Database().Collation()

	matched, err := regexp.MatchString("bin+|ci+", collation)
	if err != nil {
		t.Errorf("Regex error while matching collation string\n%v", err)
	}
	if matched {
		return
	}
	t.Error("Collation is not valid")
}

func TestColumn(t *testing.T) {
	col := f.Database().Column()
	if col != "" {
		return
	}
	t.Error("Could not generate database column")
}

func TestType(t *testing.T) {
	dbType := f.Database().Type()
	if dbType != "" {
		return
	}
	t.Error("Could not generate database type")
}
func TestEngine(t *testing.T) {
	eng := f.Database().Engine()
	if eng != "" {
		return
	}
	t.Error("Could not generate database engine")
}
