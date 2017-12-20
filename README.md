#### Faker

Go library that generates fake data. A Golang port of [faker.js](https://github.com/Marak/faker.js).

This is still in active development.

## Usage Example:

```go
package main

import (
	"fmt"
	"github.com/Arinzeokek/faker"
)

 func main() {
     // Get faker instance setting default locale to en (english)
     f, err := faker.New(&Config{defaultLocale: "en"})
     if err != nil {
         fmt.Errorf("Error loading locale, %v", err)
     }

     // print fake first name
     fmt.Println(f.Name().FirstName())

     // print fake job title
     fmt.Println(f.Name().JobTitle())

     //change locale(language) to example: german. (de)
     f.Lang("de")

     // print fake German full name
     fmt.Println(f.Name().FullName())

     //change locale(language) back to english
     f.Lang("en")

     // print fake English full name
     fmt.Println(f.Name().FullName())
 }
```

## API

* Address
  * ZipCode
  * City
  * CityPrefix
  * CitySuffix
  * StreetName
  * StreetAddress
  * StreetSuffix
  * StreetPrefix
  * SecondaryAddress
  * County
  * Country
  * CountryCode
  * State
  * StateAbbr
  * Latitude
  * Longitude
* Name
  * FirstName
  * LastName
  * FindName
  * JobTitle
  * Prefix
  * Suffix
  * Title
  * JobDescriptor
  * JobArea
  * JobType
* Commerce (TODO)
  * Color
  * Department
  * ProductName
  * Price
  * ProductAdjective
  * ProductMaterial
  * Product
* Company (TODO)
  * Suffixes
  * CompanyName
  * CompanySuffix
  * CatchPhrase
  * Bs
  * CatchPhraseAdjective
  * CatchPhraseDescriptor
  * CatchPhraseNoun
  * BsAdjective
  * BsBuzz
  * BsNoun
* Database (TODO)
  * Column
  * Type
  * Collation
  * Engine
* Date (TODO)
  * Past
  * Future
  * Between
  * Recent
  * Soon
  * Month
  * Weekday
* Finance (TODO)
  * Account
  * AccountName
  * Mask
  * Amount
  * TransactionType
  * CurrencyCode
  * CurrencyName
  * CurrencySymbol
  * BitcoinAddress
  * EthereumAddress
  * Iban
  * Bic
* Hacker (TODO)
  * Abbreviation
  * Adjective
  * Noun
  * Verb
  * Ingverb
  * Phrase
* Image (TODO)
  * Image
  * Avatar
  * ImageUrl
  * Abstract
  * Animals
  * Business
  * Cats
  * City
  * Food
  * Nightlife
  * Fashion
  * People
  * Nature
  * Sports
  * Technics
  * Transport
  * DataUri
* Internet (TODO)
  * Avatar
  * Email
  * ExampleEmail
  * UserName
  * Protocol
  * Url
  * DomainName
  * DomainSuffix
  * DomainWord
  * Ip
  * Ipv6
  * UserAgent
  * Color
  * Mac
  * Password
* Lorem (TODO)
  * Word
  * Words
  * Sentence
  * Slug
  * Sentences
  * Paragraph
  * Paragraphs
  * Text
  * Lines
* Phone (TODO)
  * PhoneNumber
  * PhoneNumberFormat
  * PhoneFormats
* Random (TODO)
  * Number
  * ArrayElement
  * ObjectElement
  * Uuid
  * Boolean
  * Word
  * Words
  * Image
  * Locale
  * AlphaNumeric
  * HexaDecimal
* System (TODO)
  * FileName
  * CommonFileName
  * MimeType
  * CommonFileType
  * CommonFileExt
  * FileType
  * FileExt
  * DirectoryPath
  * FilePath
  * Semver

## Road Map

* Finish API Implementation
* Compulsory Unique Value Feature Implementation
* Write tests
* Better Documentation

## Credit

* Data extracted from [faker.js](https://github.com/Marak/faker.js)

## License

* MIT
