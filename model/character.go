package model

import "net/url"

type Character struct {
	ID      string
	PC      bool
	Picture url.URL
	Name    string
}
