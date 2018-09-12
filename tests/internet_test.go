package main

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/olucode/faker"
)

var f, err = faker.New(&faker.Config{DefaultLocale: "en"}, "../locale")

func TestAvatar(t *testing.T) {
	// It should return a URL of the format
	// https://s3.amazonaws.com/uifaces/faces/twitter/*
	if err != nil {
		fmt.Println(err)
	}
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
	fmt.Printf("Color hex is %v\n", color)
	matched, err := regexp.MatchString("^#[A-F0-9]{6}$", color)
	if err != nil {
		t.Errorf("Regex error while matching color hex string\n%v", err)
	}
	if matched {
		return
	}
	fmt.Println(matched)
	t.Error("Color Hex is not valid")
}
