package mydict

import "errors"

type Dictionary map[string]string

var errNotFound =  errors.New("Not Found")
var errWordExists = errors.New("단어가 이미 존재합니다.")
var errNotExists = errors.New("단어가 존재하지 않습니다.")

func (d Dictionary) Search(word string) (string, error) {

	value, exist := d[word]

	if exist {
		return value, nil
	} else {
		return "", errNotFound
	}
}

func (d Dictionary) Add(word, def string) error {

	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {

	_, err := d.Search(word)

	switch err {
	case errWordExists:
		return errNotExists
	case nil:
		d[word] = def
	}

	return nil
}

func (d Dictionary) Delete(word string) error {

	_, err := d.Search(word)

	switch err {
	case errNotFound:
		return errNotExists
	case nil:
		delete(d, word)
	}

	return nil
}