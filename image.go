package faker

import (
	"strconv"
)

// Imager interface
type Imager interface {
	Avatar() string
	ImageURL(width, height int, category string) string
}

// Image struct
type Image struct {
	*Fake
}

// Avatar Generates a URL for an avatar.
func (i *Image) Avatar() string {
	return i.Internet().Avatar()
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
