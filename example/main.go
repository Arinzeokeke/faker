package main

import (
	"fmt"

	"github.com/olucode/faker"
)

func main() {
	f, err := faker.New(&faker.Config{DefaultLocale: "en"}, "locale")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f.Name().FirstName())
	fmt.Println(f.Name().JobTitle())
	fmt.Println(f.Internet().Avatar())
	fmt.Println(f.Internet().Mac(":"))
	fmt.Println(f.Internet().Email("Olucode", "falomo", "xmail.com"))
	fmt.Println(f.Name().FullName())
	fmt.Println(f.Lang("de").Name().Suffix())
	fmt.Println(f.Name().FullName())

}
