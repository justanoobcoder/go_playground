package dictionary

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound  = DictionaryErr("could not find the word you looking for")
	ErrWordExits = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dic Dictionary, key string) (string, error) {
	result, ok := dic[key]
	if !ok {
		return "", ErrNotFound
	}
	return result, nil
}

func (dic Dictionary) Add(key, value string) error {
	_, err := Search(dic, key)
	switch err {
	case ErrNotFound:
		dic[key] = value
	case nil:
		return ErrWordExits
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Update(key, newValue string) error {
	_, err := Search(dic, key)
	switch err {
	case ErrNotFound:
		return ErrNotFound
	case nil:
		dic[key] = newValue
	default:
		return err
	}
	return nil
}

func (dic Dictionary) Delete(key string) {
	delete(dic, key)
}
