package maps

var (
	ErrorNotFound     = DictionaryErr("could not find the word")
	ErrorWordExist    = DictionaryErr("cannot add word because it already exists")
	ErrorWordNotExist = DictionaryErr("cannot update word because it not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	res, ok := d[word]
	if !ok {
		return "", ErrorNotFound
	}
	return res, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrorNotFound:
		return ErrorWordNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
