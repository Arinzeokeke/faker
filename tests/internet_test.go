package main

import (
	"regexp"
	"strings"
	"testing"

	"github.com/olucode/faker"
)

type strslice []string

var f, err = faker.New(&faker.Config{DefaultLocale: "en"}, "../locale")

func init() {
	if err != nil {
		panic(err)
	}
}

func TestAvatar(t *testing.T) {
	// It should return a URL of the format
	// https://s3.amazonaws.com/uifaces/faces/twitter/*
	avatar := f.Internet().Avatar()

	matched, err := regexp.MatchString("https://s3.amazonaws.com/uifaces/faces/twitter/*", avatar)
	if err != nil {
		t.Errorf("Regex error while matching string\n%v", err)
	}
	if matched {
		return
	}

	t.Error("Avatar URL is not of the correct format")
}

func TestColor(t *testing.T) {
	color := f.Internet().Color(100, 100, 100)

	matched, err := regexp.MatchString("^#[A-F0-9]{6}$", color)
	if err != nil {
		t.Errorf("Regex error while matching color hex string\n%v", err)
	}
	if matched {
		return
	}

	t.Error("Color Hex is not valid")
}

func TestIP(t *testing.T) {
	// IP address should have four parts
	ip := f.Internet().IP()

	parts := strings.Split(ip, ".")
	if len(parts) == 4 {
		return
	}

	t.Errorf("Expected IP address of 4 parts, got %v parts ", len(parts))
}

func TestIPv6(t *testing.T) {
	ipv6 := f.Internet().IPv6()

	parts := strings.Split(ipv6, ".")
	if len(parts) == 8 {
		return
	}

	t.Errorf("Expected IP address of 8 parts, got %v parts ", len(parts))
}

func TestProtocol(t *testing.T) {
	protocol := f.Internet().Protocol()
	protocols := strslice{"http", "https"}

	if protocols.indexOf(protocol) != -1 {
		return
	}

	t.Errorf("Expected protocol to be http|https. got %v", protocol)
}

func (slice strslice) indexOf(searchElement string) int {
	for index, val := range slice {
		if val == searchElement {
			return index
		}
	}
	return -1
}
