package dict

import "errors"

// map은 두개의 값을 반환, (값, 값 존재여부 bool)
type Dictionary map[string]string

var (
	errNotFound   = errors.New("not Found")
	errCantUpdate = errors.New("can not update non-existsing word")
	errCantDelete = errors.New("can not delete non-existsing word")
	errWordExists = errors.New("word already exists")
)

// method는 type에 적용 가능
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
