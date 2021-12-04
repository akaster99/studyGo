package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Word is not found")

func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	} else {
		return value, errNotFound
	}
}
