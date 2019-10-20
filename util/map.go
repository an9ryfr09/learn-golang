package util

import "errors"

//Dictionary ...
type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

//Search ...
func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

//Add ...
func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
