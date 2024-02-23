package handler

import "net/http"

type SumHandler struct{}

func (h *SumHandler) Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("sum handler"))
}

func NewSumHandler() *SumHandler {
	return &SumHandler{}
}
