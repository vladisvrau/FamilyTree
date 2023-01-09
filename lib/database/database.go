package database

import "errors"

var (
	ErrEntityNotFound   = errors.New("entity not found")
	ErrInvalidInsert    = errors.New("invalid insert")
	ErrTableDoesntExist = errors.New("table doesn't exist")
)
