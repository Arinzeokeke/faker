package faker

import (
	"regexp"
	"strconv"
	"testing"
)

func TestImageURL(t *testing.T) {
	expectedURL := "http://lorempixel.com/600/400/food?200"
	imgurl := f.Image().ImageURL(600, 400, "food")

	regex := `http://lorempixel.com/600/400/food\?`

	matched, err := regexp.MatchString(regex+"[1-999]", imgurl)
	if err != nil {
		t.Errorf("Regex error while matching string\n%v", err)
	}
	if matched {
		return
	}

	t.Errorf("Expected image url: %s, \nGot: %s", expectedURL, imgurl)
}

func TestDataURL(t *testing.T) {
	rawPrefix := `data:image/svg+xml;charset=UTF-8,`
	svgString := `<svg xmlns="http://www.w3.org/2000/svg" version="1.1" baseProfile="full" width="` + strconv.Itoa(600) + `" height="` + strconv.Itoa(400) + `"> <rect width="100%" height="100%" fill="grey"/>  <text x="0" y="20" font-size="20" text-anchor="start" fill="white">` + strconv.Itoa(600) + `x` + strconv.Itoa(400) + `</text> </svg>`
	expectedDataURL := rawPrefix + svgString

	dataURL := f.Image().DataURL(600, 400)
	if expectedDataURL == dataURL {
		return
	}

	t.Errorf("Non-matching data urls. \nExpected: \n%v,\nGot: \n%v", expectedDataURL, dataURL)
}
