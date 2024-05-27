package handler

import "github.com/TropicalDog17/alert-crud/internal/storage"

type Handler struct {
	dB storage.StorageInterface
}

func NewHandler(db storage.StorageInterface) *Handler {
	return &Handler{
		dB: db,
	}
}

func (h *Handler) GetDB() storage.StorageInterface {
	return h.dB
}
