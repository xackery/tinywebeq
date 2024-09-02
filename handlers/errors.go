package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

func (h *Handlers) serverErrorResponse(w http.ResponseWriter, err error) {
	h.logger.Error("error in handler", zap.Error(err))
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
