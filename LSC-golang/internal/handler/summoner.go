package handler

import "net/http"

type SummonerHandler struct {}

func NewSummonerHandler() *SummonerHandler {
    return &SummonerHandler{}
}

func (h *SummonerHandler) GetByName(w http.ResponseWriter, r *http.Request) {
    _ = h
    _ = r
    http.Error(w, "not implemented", http.StatusNotImplemented)
}
