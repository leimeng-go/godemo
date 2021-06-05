package error

import (
	"database/sql"
	"errors"
	"fmt"
)

func NewCustomError(content string) error {
	return errors.New(content)
}
func NewCustomError1() error {
	new := sql.ErrNoRows
	return fmt.Errorf("out: haha,inner : %w", new)
}

type MyError struct {
	err error
	msg string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%s", e.err.Error(), e.msg)
}
