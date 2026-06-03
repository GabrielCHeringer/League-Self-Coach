package handler

import "net/http"

type StatsHandler struct {}

func NewStatsHandler() *StatsHandler {
    return &StatsHandler{}
}

func (h *StatsHandler) GetByPUUID(w http.ResponseWriter, r *http.Request) {
    _ = h
    _ = r
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
