package handler

import "github.com/rickluonz/pawsitive/cmd/api/db"

type Handler struct {
	DB *db.DB
}

func New(db *db.DB) *Handler {
	return &Handler{DB: db}
}