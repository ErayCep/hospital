package handlers

import (
	"hospital/db"
)

type Handlers struct {
	Storage db.Storage
}

func NewHandler(storage db.Storage) *Handlers {
	return &Handlers{
		Storage: storage,
	}
}
