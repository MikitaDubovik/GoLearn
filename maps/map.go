package maps

import "errors"

var (
	ErrNotFound        = errors.New("could not find the key you were looking for")
	ErrKeyExists       = errors.New("key already exists")
	ErrKeyDoesNotExist = errors.New("key does not exist")
)

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(newKey, newValue string) error {
	_, err := d.Search(newKey)
	switch err {
	case ErrNotFound:
		d[newKey] = newValue
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrKeyDoesNotExist
	case nil:
		delete(d, key)
	default:
		return err
	}

	return err
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}
