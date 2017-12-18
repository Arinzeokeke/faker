package faker

import (
	"fmt"
	"os"

	"github.com/Arinzeokeke/tr"
)

// Fake Object
type Fake struct {
	locale        string
	defaultLocale string
}

type Faker interface {
	Name() Namer
}

func init() {
	engine, err := tr.Init("locale", "en")
	if err != nil {
		fmt.Printf("Couldn't load locale, Error: %v", err)
		os.Exit(1)

	}

	fmt.Println(engine.Tr("address"))

}
