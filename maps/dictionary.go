package maps

import "errors"

type Dictionary map[string]string

var ErrWordNotFound = errors.New("word not found")

func (d Dictionary) Search(word string) (string, error) {
	definition, exist := d[word]
	if !exist {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) {
	d[key] = value
}
