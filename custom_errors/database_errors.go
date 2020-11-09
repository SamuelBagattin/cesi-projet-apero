package custom_errors

import "fmt"

type EntityNotFound struct {
	Id int
}

func (e EntityNotFound) Error() string {
	return fmt.Sprintf("Entity not found : %d", e.Id)
}
