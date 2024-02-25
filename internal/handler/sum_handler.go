package handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

type SumService interface {
	Sum(a, b int) (int, error)
}

type SumHandler struct {
	sumService SumService
}

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
	slog.Info(fmt.Sprintf("Request: %v", req))

	res, err := h.sumService.Sum(req.A, req.B)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("res: %v", res)))
}

func NewSumHandler(sumService SumService) *SumHandler {
	return &SumHandler{
		sumService: sumService,
	}
}
