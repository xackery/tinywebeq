package handlers

import "net/http"

func (h *Handlers) Root() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("Hello, world!")); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
