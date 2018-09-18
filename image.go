package faker

import (
	"fmt"
	"strconv"
)

// Imager interface
type Imager interface {
	Avatar() string
	RandImage(width, height int) string
	ImageURL(width, height int, category string) string
	DataURL(width, height int) string
	Abstract(width, height int) string
	Animals(width, height int) string
	Business(width, height int) string
	Cats(width, height int) string
	City(width, height int) string
	Food(width, height int) string
	Nightlife(width, height int) string
	Fashion(width, height int) string
	People(width, height int) string
	Nature(width, height int) string
	Sports(width, height int) string
	Technics(width, height int) string
	Transport(width, height int) string
}

// Image struct
type Image struct {
	*Fake
}

// Avatar Generates a URL for an avatar.
func (i *Image) Avatar() string {
	return i.Internet().Avatar()
}

// RandImage Returns a random image url from any of the available categories
func (i *Image) RandImage(width, height int) string {
	categories := []string{
		"abstract", "animals", "business", "cats", "city",
		"food", "nightlife", "fashion", "people",
		"nature", "sports", "technics", "transport",
	}
	categoriesFuncs := map[string]interface{}{
		"abstract":  i.Image().Abstract,
		"animals":   i.Image().Animals,
		"business":  i.Image().Business,
		"cats":      i.Image().Cats,
		"city":      i.Image().City,
		"food":      i.Image().Food,
		"nightlife": i.Image().Nightlife,
		"fashion":   i.Image().Fashion,
		"people":    i.Image().People,
		"nature":    i.Image().Nature,
		"sports":    i.Image().Sports,
		"technics":  i.Image().Technics,
		"transport": i.Image().Transport,
	}

	function := selectElement(categories)

	result, err := callFunc(categoriesFuncs, function, width, height)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", result[0])
}

// ImageURL Generates an image url
func (i *Image) ImageURL(width, height int, category string) string {
	protocol := "http://"
	w := strconv.Itoa(width)
	h := strconv.Itoa(height)

	url := protocol + "lorempixel.com/" + w + "/" + h + "/" + category

	url = url + "?" + strconv.Itoa(random(999))

	return url
}

// DataURL Returns a dataurl format
func (i *Image) DataURL(width, height int) string {
	rawPrefix := `data:image/svg+xml;charset=UTF-8,`
	svgString := `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" baseProfile="full" width="` + strconv.Itoa(width) + `" height="` + strconv.Itoa(height) + `"> <rect width="100%" height="100%" fill="grey"/>  <text x="0" y="20" font-size="20" text-anchor="start" fill="white">` + strconv.Itoa(width) + `x` + strconv.Itoa(height) + `</text> </svg>`
	return rawPrefix + svgString
}

// Abstract Returns an image in the Abstract category
func (i *Image) Abstract(width, height int) string {
	return i.Image().ImageURL(width, height, "abstract")
}

// Animals Returns an image in the Animals category
func (i *Image) Animals(width, height int) string {
	return i.Image().ImageURL(width, height, "animals")
}

// Business Returns an image in the Business category
func (i *Image) Business(width, height int) string {
	return i.Image().ImageURL(width, height, "business")
}

// Cats Returns an image in the Cats category
func (i *Image) Cats(width, height int) string {
	return i.Image().ImageURL(width, height, "cats")
}

// City Returns an image in the City category
func (i *Image) City(width, height int) string {
	return i.Image().ImageURL(width, height, "city")
}

// Food Returns an image in the Food category
func (i *Image) Food(width, height int) string {
	return i.Image().ImageURL(width, height, "food")
}

// Nightlife Returns an image in the Nightlife category
func (i *Image) Nightlife(width, height int) string {
	return i.Image().ImageURL(width, height, "nightlife")
}

// Fashion Returns an image in the Fashion category
func (i *Image) Fashion(width, height int) string {
	return i.Image().ImageURL(width, height, "fashion")
}

// People Returns an image in the People category
func (i *Image) People(width, height int) string {
	return i.Image().ImageURL(width, height, "people")
}

// Nature Returns an image in the Nature category
func (i *Image) Nature(width, height int) string {
	return i.Image().ImageURL(width, height, "nature")
}

// Sports Returns an image in the Sports category
func (i *Image) Sports(width, height int) string {
	return i.Image().ImageURL(width, height, "sports")
}

// Technics Returns an image in the Technics category
func (i *Image) Technics(width, height int) string {
	return i.Image().ImageURL(width, height, "technics")
}

// Transport Returns an image in the Transport category
func (i *Image) Transport(width, height int) string {
	return i.Image().ImageURL(width, height, "transport")
}
