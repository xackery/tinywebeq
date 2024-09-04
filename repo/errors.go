package repo

import "errors"

var ErrNotFound = errors.New("not found")
var ErrNotPointer = errors.New("target must be a pointer to a struct")
