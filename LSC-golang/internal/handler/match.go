package handler

import "net/http"

type MatchHandler struct {}

func NewMatchHandler() *MatchHandler {
    return &MatchHandler{}
}

func (h *MatchHandler) GetByPUUID(w http.ResponseWriter, r *http.Request) {
    _ = h
    _ = r
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
