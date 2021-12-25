package handler

import "http-callback/server/usecase"

type Handler struct {
	UC *usecase.UC
}

func New(uc *usecase.UC) *Handler { return &Handler{UC: uc} }
