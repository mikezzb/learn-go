package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")
var ErrWordExists = errors.New("word exists")
var ErrWordDNE = errors.New("Word DNE")

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, ok := d[key]
	if ok {
		return ErrWordExists
	}
	d[key] = val
	return nil
}

func (d Dictionary) Update(key, val string) error {
	_, ok := d[key]
	if !ok {
		return ErrWordDNE
	}
	d[key] = val
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, ok := d[key]
	if !ok {
		return ErrWordDNE
	}
	delete(d, key)
	return nil
}
