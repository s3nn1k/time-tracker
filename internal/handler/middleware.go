package handler

import (
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

type nextHandler func(w http.ResponseWriter, r *http.Request, id int)

func (h *Handler) Authorize(handler nextHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get(authorizationHeader)
		if authHeader == "" {
			http.Error(w, "authorization header empty", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		id, err := h.Auth.ParseToken(headerParts[1])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}

		handler(w, r, id)
	})
}
