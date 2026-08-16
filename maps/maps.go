package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrKeyExsited = DictionaryErr("the key has exsited")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	if difinition, ok := d[word]; ok {
		return difinition, nil
	}
	return "", ErrNotFound
}

// 仅加入键值对，不覆盖; 如果key已存在则返回ErrKeyExsited
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		return ErrKeyExsited
	case ErrNotFound:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if _, ok := d[word]; !ok {
		return ErrNotFound
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
