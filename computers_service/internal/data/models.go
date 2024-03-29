package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Computers ComputerModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Computers: ComputerModel{DB: db},
	}
}
