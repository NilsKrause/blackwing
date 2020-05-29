package model

import (
	"errors"
	"net/url"
	"strings"
)

// Character is a the structure that holds a pc or a npc
type Character struct {
	ID      string
	PC      bool
	Picture *url.URL
	Name    string
}

// ErrNotPictureURL gets returned when the url is not a suppored type
var ErrNotPictureURL = errors.New("urls must point to a png or jpg")

// NewCharacter returns a new initialized character
func NewCharacter(isPc bool, surl string, name string) (*Character, error) {
	// generate id\

	if !strings.HasSuffix(surl, ".jpg") ||
		!strings.HasSuffix(surl, ".png") {
		return nil, ErrNotPictureURL
	}

	u, err := url.Parse(surl)
	if err != nil {
		return nil, err
	}

	return &Character{PC: isPc, Name: name, Picture: u}, nil

}
