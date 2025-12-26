package errs

import "errors"

var (
	ErrNotFound     = errors.New("Todo not found")
	ValidationError = errors.New("Title is required")
)
