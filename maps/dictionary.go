package main

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

var ErrNilMap = errors.New("map is nil")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(word, definition string) error {
	if d == nil {
		return ErrNilMap
	}
	d[word] = definition
	return nil
}
