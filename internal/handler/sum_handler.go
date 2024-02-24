package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type SumHandler struct{}

type SumRequest struct {
	A int
	B int
}

func (h *SumHandler) Handler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Request: %s", r.Body)

	req := SumRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error("%s", err)
		return
	}
	slog.Info("Request: %+v", req)

	w.Write([]byte("sum handler"))
}

func NewSumHandler() *SumHandler {
	return &SumHandler{}
}
