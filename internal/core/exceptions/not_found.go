package exceptions

import "fmt"

type errNotFound struct {
	base string
	id   interface{}
}

func (e *errNotFound) Error() string {
	return fmt.Sprintf("a %s with id %s not found", e.base, e.id)
}

func EntityNotFound(base string, id interface{}) error {
	return &errNotFound{
		base: base,
		id:   id,
	}
}
