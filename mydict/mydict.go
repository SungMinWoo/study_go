package mydict

type Dictionary map[string]string

var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWorkExists = errors.New("That work already exists")
)

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
		return errWorkExists
	}
	return nil
	// 위에 if문을 switch문으로 작성할 경우
	// switch err {
	// case errNotFound:
	// 	d[word] = def
	// case nil:
	// 	return errWorkExists
	// }
	// return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}


func (d Dictionary) Delete(word string) {
	delete(d, word)
}